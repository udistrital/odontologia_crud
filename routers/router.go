// @APIVersion 1.0
// @Title Odontologia
// @Description Api para Odontología dentro del módulo de Salud del proyecto SIBUD
// @Contact baluist@correo.udistrital.edu.co
// @TermsOfServiceUrl http://www.udistrital.edu.co/politicas-de-privacidad.pdf
// @License BSD-3-Clause
// @LicenseUrl http://opensource.org/licenses/BSD-3-Clause
// @BasePath /Odontologia/v1
package routers

import (
	"Odontologia/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/Anamnesis",
			beego.NSInclude(
				&controllers.AnamnesisController{},
			),
		),
		beego.NSNamespace("/Diagnostico",
			beego.NSInclude(
				&controllers.DiagnosticoController{},
			),
		),
		beego.NSNamespace("/ExamenDental",
			beego.NSInclude(
				&controllers.ExamenDentalController{},
			),
		),
		beego.NSNamespace("/ExamenesComplementarios",
			beego.NSInclude(
				&controllers.ExamenesComplementariosController{},
			),
		),
		beego.NSNamespace("/ExamenEstomatologico",
			beego.NSInclude(
				&controllers.ExamenEstomatologicoController{},
			),
		),
		beego.NSNamespace("/Odontograma",
			beego.NSInclude(
				&controllers.OdontogramaController{},
			),
		),
		beego.NSNamespace("/TipoOdontograma",
			beego.NSInclude(
				&controllers.TipoOdontogramaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
