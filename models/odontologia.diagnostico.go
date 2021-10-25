package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type OdontologiaDiagnostico struct {
	IdDiagnostico  int                   `orm:"column(id_diagnostico);pk;auto"`
	Diagnostivo    string                `orm:"column(diagnostivo);null"`
	Pronostico     string                `orm:"column(pronostico);null"`
	Evolucion      string                `orm:"column(evolucion);null"`
	Observaciones  string                `orm:"column(observaciones);null"`
	IdHojaHistoria *MedicinaHojaHistoria `orm:"column(id_hoja_historia);rel(fk);null"`
}

func (t *OdontologiaDiagnostico) TableName() string {
	return "diagnosticoodontologia"
}
func init() {
	orm.RegisterModel(new(OdontologiaDiagnostico))
}

// AddOdontologiaDiagnostico inserta un registro en la tabla diagnosticoodontologia
// Último registro insertado con éxito
func AddOdontologiaDiagnostico(m *OdontologiaDiagnostico) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOdontologiaDiagnosticoById obtiene un registro de la tabla diagnosticoodontologia por su id
// Id no existe
func GetOdontologiaDiagnosticoById(id int) (v *OdontologiaDiagnostico, err error) {
	o := orm.NewOrm()
	v = &OdontologiaDiagnostico{IdDiagnostico: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOdontologiaDiagnostico obtiene todos los registros de la tabla diagnosticoodontologia
// No existen registros
func GetAllOdontologiaDiagnostico(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(OdontologiaDiagnostico))
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
	var l []OdontologiaDiagnostico
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

// UpdateOdontologiaDiagnostico actualiza un registro de la tabla diagnosticoodontologia
// El registro a actualizar no existe
func UpdateOdontologiaDiagnostico(m *OdontologiaDiagnostico) (err error) {
	o := orm.NewOrm()
	v := OdontologiaDiagnostico{IdDiagnostico: m.IdDiagnostico}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Numero de registros actualizados:", num)
		}
	}
	return
}

// DeleteOdontologiaDiagnostico elimina un registro de la tabla diagnosticoodontologia
// El registro a eliminar no existe
func DeleteOdontologiaDiagnostico(id int) (err error) {
	o := orm.NewOrm()
	v := OdontologiaDiagnostico{IdDiagnostico: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&OdontologiaDiagnostico{IdDiagnostico: id}); err == nil {
			fmt.Println("Numero de registros eliminados:", num)
		}
	}
	return
}
