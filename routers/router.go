// @APIVersion 1.0.0
// @Title beego CMDB API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"ss_cmdb/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1/conf",
		/*
			beego.NSNamespace("/object",
				beego.NSInclude(
					&controllers.ObjectController{},
				),
			),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		*/
		beego.NSNamespace("/ini",
			beego.NSInclude(
				&controllers.Info_iniController{},
			),
		),
	)
	beego.AddNamespace(ns)
}