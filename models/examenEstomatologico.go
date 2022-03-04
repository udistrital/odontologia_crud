package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ExamenEstomatologico struct {
	Id                   		int    `orm:"column(id_examen_estomatologico);pk;auto"`
	HistoriaClinicaId    		int    `orm:"column(id_historia_clinica);null"`
	HojaHistoriaId       		int    `orm:"column(id_hoja_historia);null"`
	ArticulacionTemporo  		string `orm:"column(articulacion_temporo);null"`
	Labios               		string `orm:"column(labios);null"`
	Lengua               		string `orm:"column(lengua);null"`
	Paladar              		string `orm:"column(paladar);null"`
	PisoBoca             		string `orm:"column(piso_boca);null"`
	Carrillos            		string `orm:"column(carrillos);null"`
	GlandulasSalivares   		string `orm:"column(glandulas_salivares);null"`
	Maxilares            		string `orm:"column(maxilares);null"`
	SenosMaxilares       		string `orm:"column(senos_maxilares);null"`
	MusculosMasticadores 		string `orm:"column(musculos_masticadores);null"`
	SistemaVascular      		string `orm:"column(sistema_vascular);null"`
	SistemaNervioso      		string `orm:"column(sistema_nervioso);null"`
	SistemaLinfaticoRegional 	string `orm:"column(sistema_linfatico_regional);null"`
	FechaCreacion     *time.Time `orm:"column(fecha_creacion);type(timestamp without time zone);null"`
	FechaModificacion *time.Time `orm:"column(fecha_modificacion);type(timestamp without time zone);null"`
	Activo            bool       `orm:"column(activo);null"`
}

func (t *ExamenEstomatologico) TableName() string {
	return "examenestomatologico"
}
func init() {
	orm.RegisterModel(new(ExamenEstomatologico))
}

// AddExamenEstomatologico inserta un registro en la tabla examenestomatologico
// Último registro insertado con éxito
func AddExamenEstomatologico(m *ExamenEstomatologico) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetExamenEstomatologicoById obtiene un registro de la tabla examenestomatologico por su id
// Id no existe
func GetExamenEstomatologicoById(id int) (v *ExamenEstomatologico, err error) {
	o := orm.NewOrm()
	v = &ExamenEstomatologico{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllExamenEstomatologico obtiene todos los registros de la tabla examenestomatologico
// No existen registros
func GetAllExamenEstomatologico(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ExamenEstomatologico))
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
	var l []ExamenEstomatologico
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

// UpdateExamenEstomatologico actualiza un registro de la tabla examenestomatologico
// El registro a actualizar no existe
func UpdateExamenEstomatologico(m *ExamenEstomatologico) (err error) {
	o := orm.NewOrm()
	v := ExamenEstomatologico{Id: m.Id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Numero de registros actualizados:", num)
		}
	}
	return
}

// DeleteExamenEstomatologico elimina un registro de la tabla examenestomatologico
// El registro a eliminar no existe
func DeleteExamenEstomatologico(id int) (err error) {
	o := orm.NewOrm()
	v := ExamenEstomatologico{Id: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ExamenEstomatologico{Id: id}); err == nil {
			fmt.Println("Numero de registros eliminados:", num)
		}
	}
	return
}
