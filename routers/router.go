// @APIVersion 1.0.0
// @Title data-service API
// @Description data-service only serve summary
// @Contact qsg@corex-tek.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"app-service/data-service/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/data",
			beego.NSInclude(
				&controllers.DataController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
