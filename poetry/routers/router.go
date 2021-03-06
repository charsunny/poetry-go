// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"poetry/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/common",
			beego.NSInclude(
				&controllers.BaseController{},
			),
		),
		beego.NSNamespace("/rec",
			beego.NSInclude(
				&controllers.RecommandController{},
			),
		),
		beego.NSNamespace("/poem",
			beego.NSInclude(
				&controllers.PoemController{},
			),
		),
		beego.NSNamespace("/feed",
			beego.NSInclude(
				&controllers.FeedController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/col",
			beego.NSInclude(
				&controllers.ColumnController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
