// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
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
		beego.NSNamespace("/ControlPlaca",
			beego.NSInclude(
				&controllers.ControlPlacaController{},
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
