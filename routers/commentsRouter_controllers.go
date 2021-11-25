package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {
	beego.GlobalControllerRouter["Odontologia/controllers:AnamnesisController"] = append(beego.GlobalControllerRouter["Medicina/controllers:AnamnesisController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:AnamnesisController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:AnamnesisController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:AnamnesisController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:AnamnesisController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:AnamnesisController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:AnamnesisController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:AnamnesisController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:AnamnesisController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:DiagnosticoController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:DiagnosticoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:DiagnosticoController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:DiagnosticoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:DiagnosticoController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:DiagnosticoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:DiagnosticoController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:DiagnosticoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:DiagnosticoController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:DiagnosticoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:ExamenDentalController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:ExamenDentalController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:ExamenDentalController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:ExamenDentalController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:ExamenDentalController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:ExamenDentalController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:ExamenDentalController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:ExamenDentalController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:ExamenDentalController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:ExamenDentalController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:ExamenEstomatologicoController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:ExamenEstomatologicoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:ExamenEstomatologicoController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:ExamenEstomatologicoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:ExamenEstomatologicoController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:ExamenEstomatologicoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:ExamenEstomatologicoController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:ExamenEstomatologicoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:ExamenEstomatologicoController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:ExamenEstomatologicoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:ExamenesComplementariosController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:ExamenesComplementariosController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:ExamenesComplementariosController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:ExamenesComplementariosController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:ExamenesComplementariosController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:ExamenesComplementariosController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:ExamenesComplementariosController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:ExamenesComplementariosController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:ExamenesComplementariosController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:ExamenesComplementariosController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:OdontogramaController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:OdontogramaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:OdontogramaController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:OdontogramaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:OdontogramaController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:OdontogramaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:OdontogramaController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:OdontogramaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:OdontogramaController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:OdontogramaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:TipoOdontogramaController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:TipoOdontogramaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:TipoOdontogramaController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:TipoOdontogramaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:TipoOdontogramaController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:TipoOdontogramaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:TipoOdontogramaController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:TipoOdontogramaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Odontologia/controllers:TipoOdontogramaController"] = append(beego.GlobalControllerRouter["Odontologia/controllers:TipoOdontogramaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
