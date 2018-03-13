package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["axell_first_rest/controllers:UserController"] = append(beego.GlobalControllerRouter["axell_first_rest/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})
	beego.GlobalControllerRouter["axell_first_rest/controllers:UserController"] = append(beego.GlobalControllerRouter["axell_first_rest/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetOneUser",
			Router: `/:userId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})
	beego.GlobalControllerRouter["axell_first_rest/controllers:UserController"] = append(beego.GlobalControllerRouter["axell_first_rest/controllers:UserController"],
		beego.ControllerComments{
			Method: "PostNewUser",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})
	beego.GlobalControllerRouter["axell_first_rest/controllers:UserController"] = append(beego.GlobalControllerRouter["axell_first_rest/controllers:UserController"],
		beego.ControllerComments{
			Method: "UpdateExistingUser",
			Router: `/:userId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})
	beego.GlobalControllerRouter["axell_first_rest/controllers:UserController"] = append(beego.GlobalControllerRouter["axell_first_rest/controllers:UserController"],
		beego.ControllerComments{
			Method: "DeleteUser",
			Router: `/:userId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})
}
