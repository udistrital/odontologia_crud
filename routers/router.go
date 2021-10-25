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
				&controllers.OdontologiaAnamnesisController{},
			),
		),
		beego.NSNamespace("/ControlPlaca",
			beego.NSInclude(
				&controllers.OdontologiaControlPlacaController{},
			),
		),
		beego.NSNamespace("/Diagnostico",
			beego.NSInclude(
				&controllers.OdontologiaDiagnosticoController{},
			),
		),
		beego.NSNamespace("/ExamenDental",
			beego.NSInclude(
				&controllers.OdontologiaExamenDentalController{},
			),
		),
		beego.NSNamespace("/ExamenesComplementarios",
			beego.NSInclude(
				&controllers.OdontologiaExamenesComplementariosController{},
			),
		),
		beego.NSNamespace("/ExamenEstomatologico",
			beego.NSInclude(
				&controllers.OdontologiaExamenEstomatologicoController{},
			),
		),
		beego.NSNamespace("/Odontograma",
			beego.NSInclude(
				&controllers.OdontologiaOdontogramaController{},
			),
		),
		beego.NSNamespace("/TipoOdontograma",
			beego.NSInclude(
				&controllers.OdontologiaTipoOdontogramaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
