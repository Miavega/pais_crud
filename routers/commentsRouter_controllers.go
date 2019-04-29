package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/miavega/pais_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/miavega/pais_crud/controllers:CiudadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/pais_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/miavega/pais_crud/controllers:CiudadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/pais_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/miavega/pais_crud/controllers:CiudadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/pais_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/miavega/pais_crud/controllers:CiudadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/pais_crud/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/miavega/pais_crud/controllers:CiudadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/pais_crud/controllers:PaisController"] = append(beego.GlobalControllerRouter["github.com/miavega/pais_crud/controllers:PaisController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/pais_crud/controllers:PaisController"] = append(beego.GlobalControllerRouter["github.com/miavega/pais_crud/controllers:PaisController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/pais_crud/controllers:PaisController"] = append(beego.GlobalControllerRouter["github.com/miavega/pais_crud/controllers:PaisController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/pais_crud/controllers:PaisController"] = append(beego.GlobalControllerRouter["github.com/miavega/pais_crud/controllers:PaisController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/miavega/pais_crud/controllers:PaisController"] = append(beego.GlobalControllerRouter["github.com/miavega/pais_crud/controllers:PaisController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
