package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type OdontologiaOdontograma struct {
	IdOdontograma     int                         `orm:"column(id_odontograma);pk;auto"`
	IdHojaHistoria    *MedicinaHojaHistoria       `orm:"column(id_hoja_historia);rel(fk);null"`
	Diagrama          string                      `orm:"column(diagrama);null"`
	Observaciones     string                      `orm:"column(observaciones);null"`
	IdTipoOdontograma *OdontologiaTipoOdontograma `orm:"column(id_tipo_odontograma);rel(fk);null"`
}

func (p *OdontologiaOdontograma) TableName() string {
	return "odontograma"
}
func init() {
	orm.RegisterModel(new(OdontologiaOdontograma))
}

// AddOdontologiaOdontograma inserta un registro en la tabla odontograma
// Último registro insertado con éxito
func AddOdontologiaOdontograma(m *OdontologiaOdontograma) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOdontologiaOdontogramaById obtiene un registro de la tabla odontograma por su id
// Id no existe
func GetOdontologiaOdontogramaById(id int) (v *OdontologiaOdontograma, err error) {
	o := orm.NewOrm()
	v = &OdontologiaOdontograma{IdOdontograma: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOdontologiaOdontograma obtiene todos los registros de la tabla odontograma
// No existen registros
func GetAllOdontologiaOdontograma(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(OdontologiaOdontograma))
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
	var l []OdontologiaOdontograma
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

// UpdateOdontologiaOdontograma actualiza un registro de la tabla odontograma
// El registro a actualizar no existe
func UpdateOdontologiaOdontograma(m *OdontologiaOdontograma) (err error) {
	o := orm.NewOrm()
	v := OdontologiaOdontograma{IdOdontograma: m.IdOdontograma}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Numero de registros actualizados:", num)
		}
	}
	return
}

// DeleteOdontologiaOdontograma elimina un registro de la tabla odontograma
// El registro a eliminar no existe
func DeleteOdontologiaOdontograma(id int) (err error) {
	o := orm.NewOrm()
	v := OdontologiaOdontograma{IdOdontograma: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&OdontologiaOdontograma{IdOdontograma: id}); err == nil {
			fmt.Println("Numero de registros eliminados:", num)
		}
	}
	return
}
