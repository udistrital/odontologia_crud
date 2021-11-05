package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ExamenesComplementarios struct {
	Id                int    `orm:"column(id_examenes_complementarios);pk;auto"`
	HistoriaClinicaId int    `orm:"column(id_historia_clinica);null"`
	HojaHistoriaId    int    `orm:"column(id_hoja_historia);null"`
	PeriapicalInicio  int    `orm:"column(periapical_inicio);null"`
	PeriapicalFinal   int    `orm:"column(periapical_final);null"`
	PanoramicaInicio  int    `orm:"column(panoramica_inicio);null"`
	PanoramicaFinal   int    `orm:"column(panoramica_final);null"`
	OtraInicio        int    `orm:"column(otra_inicio);null"`
	OtraFinal         int    `orm:"column(otra_final);null"`
	LaboratorioInicio int    `orm:"column(laboratorio_inicio);null"`
	LaboratorioFinal  int    `orm:"column(laboratorio_final);null"`
	Tp                string `orm:"column(tp);null"`
	Tpt               string `orm:"column(tpt);null"`
	Coagulacion       string `orm:"column(coagulacion);null"`
	Sangria           string `orm:"column(sangria);null"`
	Otra              string `orm:"column(otra);null"`
}

func (t *ExamenesComplementarios) TableName() string {
	return "examenescomplementarios"
}
func init() {
	orm.RegisterModel(new(ExamenesComplementarios))
}

// AddExamenesComplementarios inserta un registro en la tabla examenescomplementarios
// Último registro insertado con éxito
func AddExamenesComplementarios(m *ExamenesComplementarios) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetExamenesComplementariosById obtiene un registro de la tabla examenescomplementarios por su id
// Id no existe
func GetExamenesComplementariosById(id int) (v *ExamenesComplementarios, err error) {
	o := orm.NewOrm()
	v = &ExamenesComplementarios{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllExamenesComplementarios obtiene todos los registros de la tabla examenescomplementarios
// No existen registros
func GetAllExamenesComplementarios(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ExamenesComplementarios))
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
	var l []ExamenesComplementarios
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

// UpdateExamenesComplementarios actualiza un registro de la tabla examenescomplementarios
// El registro a actualizar no existe
func UpdateExamenesComplementarios(m *ExamenesComplementarios) (err error) {
	o := orm.NewOrm()
	v := ExamenesComplementarios{Id: m.Id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Numero de registros actualizados:", num)
		}
	}
	return
}

// DeleteExamenesComplementarios elimina un registro de la tabla examenescomplementarios
// El registro a eliminar no existe
func DeleteExamenesComplementarios(id int) (err error) {
	o := orm.NewOrm()
	v := ExamenesComplementarios{Id: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ExamenesComplementarios{Id: id}); err == nil {
			fmt.Println("Numero de registros eliminados:", num)
		}
	}
	return
}
