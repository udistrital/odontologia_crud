package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ControlPlaca struct {
	IdcontrolPlaca    int              `orm:"column(id_control_placa);pk;auto"`
	IdHojaHistoria    int              `orm:"column(id_hoja_historia);null"`
	IdTipoOdontograma *TipoOdontograma `orm:"column(id_tipo_odontograma);rel(fk);null"`
	IndiceAnterior    int              `orm:"column(indice_anterior);null"`
	IndiceActual      int              `orm:"column(indice_actual);null"`
	Fecha             time.Time        `orm:"column(fecha);type(date);null"`
	Vestibulares      string           `orm:"column(vestibulares);type(json);null"`
	Observaciones     string           `orm:"column(observaciones);null"`
}

func (t *ControlPlaca) TableName() string {
	return "controlplaca"
}
func init() {
	orm.RegisterModel(new(ControlPlaca))
}

// AddControlPlaca inserta un registro en la tabla controlplaca
// Último registro insertado con éxito
func AddControlPlaca(m *ControlPlaca) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetControlPlacaById obtiene un registro de la tabla controlplaca por su id
// Id no existe
func GetControlPlacaById(id int) (v *ControlPlaca, err error) {
	o := orm.NewOrm()
	v = &ControlPlaca{IdcontrolPlaca: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllControlPlaca obtiene todos los registros de la tabla controlplaca
// No existen registros
func GetAllControlPlaca(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ControlPlaca))
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
	var l []ControlPlaca
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

// UpdateControlPlaca actualiza un registro de la tabla controlplaca
// El registro a actualizar no existe
func UpdateControlPlaca(m *ControlPlaca) (err error) {
	o := orm.NewOrm()
	v := ControlPlaca{IdcontrolPlaca: m.IdcontrolPlaca}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Número de registros actualizados de la base de datos:", num)
		}
	}
	return
}

// DeleteControlPlaca elimina un registro de la tabla controlplaca
// El registro a eliminar no existe
func DeleteControlPlaca(id int) (err error) {
	o := orm.NewOrm()
	v := ControlPlaca{IdcontrolPlaca: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ControlPlaca{IdcontrolPlaca: id}); err == nil {
			fmt.Println("Número de registros eliminados de la base de datos:", num)
		}
	}
	return
}
