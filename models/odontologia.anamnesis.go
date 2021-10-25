package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Anamnesis struct {
	IdAnamnesis            int                      `orm:"column(id_anamnesis);pk;auto"`
	IdHistoriaClinica      *MedicinaHistoriaClinica `orm:"column(id_historia_clinica);rel(fk);null"`
	Tratamiento            string                   `orm:"column(tratamiento);null"`
	Medicamentos           string                   `orm:"column(medicamento);null"`
	Alergias               string                   `orm:"column(alergias);null"`
	Hemorragias            string                   `orm:"column(hemorragias);null"`
	Irradiaciones          string                   `orm:"column(irradiaciones);null"`
	Sinusitis              string                   `orm:"column(sinusitis);null"`
	EnfermedadRespiratoria string                   `orm:"column(enfermedad_respiratoria);null"`
	Cardiopatias           string                   `orm:"column(cardiopatias);null"`
	Diabetes               string                   `orm:"column(diabetes);null"`
	FiebreReumatica        string                   `orm:"column(fiebre_reumatica);null"`
	Hepatitis              string                   `orm:"column(hepatitis);null"`
	Hipertension           string                   `orm:"column(hipertension);null"`
	AntecedenteFamiliar    string                   `orm:"column(antecedente_familiar);null"`
	Cepillado              string                   `orm:"column(cepillado);null"`
	Ceda                   string                   `orm:"column(ceda);null"`
	Enjuague               string                   `orm:"column(enjuague);null"`
	Dulces                 string                   `orm:"column(dulces);null"`
	Fuma                   string                   `orm:"column(fuma);null"`
	Chicle                 string                   `orm:"column(chicle);null"`
	Otras                  string                   `orm:"column(otras);null"`
	UltimaVisita           time.Time                `orm:"column(ultima_visita);type(date);null"`
}

func (t *Anamnesis) TableName() string {
	return "anamnesis"
}
func init() {
	orm.RegisterModel(new(Anamnesis))
}

// AddAnamnesis inserta un registro en la tabla anamnesis
// Último registro insertado con éxito
func AddAnamnesis(m *Anamnesis) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAnamnesisById obtiene un registro de la tabla anamnesis por su id
// Id no existe
func GetAnamnesisById(id int) (v *Anamnesis, err error) {
	o := orm.NewOrm()
	v = &Anamnesis{IdAnamnesis: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAnamnesis consulta todas las anamnesis
// No existen registros
func GetAllAnamnesis(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Anamnesis))
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("error: Orden inválido, debe ser del tipo [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("error: Orden inválido, debe ser del tipo [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("error: los tamaños de 'sortby', 'order' no coinciden o el tamaño de 'order' no es 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("error: campos de 'order' no utilizados")
		}
	}
	var l []Anamnesis
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateAnamnesisById actualiza un registro de la tabla anamnesis
// El registro a actualizar no existe
func UpdateAnamnesis(m *Anamnesis) (err error) {
	o := orm.NewOrm()
	v := Anamnesis{IdAnamnesis: m.IdAnamnesis}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Número de registros actualizados en la base de datos:", num)
		}
	}
	return
}

// DeleteAnamnesis elimina un registro de la tabla anamnesis
// El registro a eliminar no existe
func DeleteAnamnesis(id int) (err error) {
	o := orm.NewOrm()
	v := Anamnesis{IdAnamnesis: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Anamnesis{IdAnamnesis: id}); err == nil {
			fmt.Println("Número de registros eliminados en la base de datos:", num)
		}
	}
	return
}
