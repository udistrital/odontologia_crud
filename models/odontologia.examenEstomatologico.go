package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type OdontologiaExamenEstomatologico struct {
	IdExamenEstomatologico int                   `orm:"column(id_examen_estomatologico);pk;auto"`
	Observaciones          string                `orm:"column(observaciones);null"`
	IdHojaHistoria         *MedicinaHojaHistoria `orm:"column(id_hoja_historia);rel(fk);null"`
	ArticulacionTemporo    string                `orm:"column(articulacion_temporo);null"`
	Labios                 string                `orm:"column(labios);null"`
	Lengua                 string                `orm:"column(lengua);null"`
	Paladar                string                `orm:"column(paladar);null"`
	PisoBoca               string                `orm:"column(piso_boca);null"`
	Carrillos              string                `orm:"column(carrillos);null"`
	GlandulasSalivares     string                `orm:"column(glandulas_salivares);null"`
	Maxilares              string                `orm:"column(maxilares);null"`
	SenosMaxilares         string                `orm:"column(senos_maxilares);null"`
	MusculosMasticadores   string                `orm:"column(musculos_masticadores);null"`
	SistemaNervioso        string                `orm:"column(sistema_nervioso);null"`
	SistemaMuscular        string                `orm:"column(sistema_muscular);null"`
	SistemaLinfatico       string                `orm:"column(sistema_linfatico);null"`
	SistemaRegional        string                `orm:"column(sistema_regional);null"`
}

func (t *OdontologiaExamenEstomatologico) TableName() string {
	return "examenestomatologico"
}
func init() {
	orm.RegisterModel(new(OdontologiaExamenEstomatologico))
}

// AddOdontologiaExamenEstomatologico inserta un registro en la tabla examenestomatologico
// Último registro insertado con éxito
func AddOdontologiaExamenEstomatologico(m *OdontologiaExamenEstomatologico) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOdontologiaExamenEstomatologicoById obtiene un registro de la tabla examenestomatologico por su id
// Id no existe
func GetOdontologiaExamenEstomatologicoById(id int) (v *OdontologiaExamenEstomatologico, err error) {
	o := orm.NewOrm()
	v = &OdontologiaExamenEstomatologico{IdExamenEstomatologico: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOdontologiaExamenEstomatologico obtiene todos los registros de la tabla examenestomatologico
// No existen registros
func GetAllOdontologiaExamenEstomatologico(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(OdontologiaExamenEstomatologico))
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
	var l []OdontologiaExamenEstomatologico
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

// UpdateOdontologiaExamenEstomatologico actualiza un registro de la tabla examenestomatologico
// El registro a actualizar no existe
func UpdateOdontologiaExamenEstomatologico(m *OdontologiaExamenEstomatologico) (err error) {
	o := orm.NewOrm()
	v := OdontologiaExamenEstomatologico{IdExamenEstomatologico: m.IdExamenEstomatologico}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Numero de registros actualizados:", num)
		}
	}
	return
}

// DeleteOdontologiaExamenEstomatologico elimina un registro de la tabla examenestomatologico
// El registro a eliminar no existe
func DeleteOdontologiaExamenEstomatologico(id int) (err error) {
	o := orm.NewOrm()
	v := OdontologiaExamenEstomatologico{IdExamenEstomatologico: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&OdontologiaExamenEstomatologico{IdExamenEstomatologico: id}); err == nil {
			fmt.Println("Numero de registros eliminados:", num)
		}
	}
	return
}
