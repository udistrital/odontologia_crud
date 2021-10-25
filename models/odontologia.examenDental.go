package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type OdontologiaExamenDental struct {
	IdExamenDental   int                   `orm:"column(id_examen_dental);pk;auto"`
	IdHojaHistoria   *MedicinaHojaHistoria `orm:"column(id_hoja_historia);rel(fk);null"`
	Supernumerarios  string                `orm:"column(supernumerarios);null"`
	Abrasion         string                `orm:"column(abrasion);null"`
	Manchas          string                `orm:"column(manchas);null"`
	PatologiaPulpar  string                `orm:"column(patologia_pulpar);null"`
	PlacaBlanda      string                `orm:"column(placa_blanda);null"`
	PlacaCalcificada string                `orm:"column(placa_calcificada);null"`
	Oclusion         string                `orm:"column(oclusiones);null"`
	Otros            string                `orm:"column(otros);null"`
	Observaciones    string                `orm:"column(observaciones);null"`
}

func (t *OdontologiaExamenDental) TableName() string {
	return "examendental"
}
func init() {
	orm.RegisterModel(new(OdontologiaExamenDental))
}

// AddOdontologiaExamenDental inserta un registro en la tabla examendental
// Último registro insertado con éxito
func AddOdontologiaExamenDental(m *OdontologiaExamenDental) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOdontologiaExamenDentalById obtiene un registro de la tabla examendental por su id
// Id no existe
func GetOdontologiaExamenDentalById(id int) (v *OdontologiaExamenDental, err error) {
	o := orm.NewOrm()
	v = &OdontologiaExamenDental{IdExamenDental: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOdontologiaExamenDental obtiene todos los registros de la tabla examendental
// No existen registros
func GetAllOdontologiaExamenDental(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(OdontologiaExamenDental))
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
	var l []OdontologiaExamenDental
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

// UpdateOdontologiaExamenDental actualiza un registro de la tabla examendental
// El registro a actualizar no existe
func UpdateOdontologiaExamenDental(m *OdontologiaExamenDental) (err error) {
	o := orm.NewOrm()
	v := OdontologiaExamenDental{IdExamenDental: m.IdExamenDental}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Numero de registros actualizados:", num)
		}
	}
	return
}

// DeleteOdontologiaExamenDental elimina un registro de la tabla examendental
// El registro a eliminar no existe
func DeleteOdontologiaExamenDental(id int) (err error) {
	o := orm.NewOrm()
	v := OdontologiaExamenDental{IdExamenDental: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&OdontologiaExamenDental{IdExamenDental: id}); err == nil {
			fmt.Println("Numero de registros eliminados:", num)
		}
	}
	return
}
