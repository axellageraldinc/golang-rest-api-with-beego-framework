package routers

import (
	"github.com/astaxie/beego"
	"axell_first_rest/controllers"
)

func init()  {
	ns := beego.NewNamespace(
		"/api",
		beego.NSNamespace(
			"/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}