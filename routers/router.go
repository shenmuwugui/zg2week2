// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"week2/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UsersController{},
			),
		),

		beego.NSNamespace("/deforment",
			beego.NSInclude(
				&controllers.DeformentController{},
			),
		),
	)
	beego.AddNamespace(ns)

	beego.Router("/del/:ids", &controllers.DeformentController{}, "get:Delmany")
	beego.Router("/delete/:ids", &controllers.UsersController{}, "patch:Delmanys")
}
