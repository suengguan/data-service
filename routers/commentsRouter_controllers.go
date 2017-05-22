package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["app-service/data-service/controllers:DataController"] = append(beego.GlobalControllerRouter["app-service/data-service/controllers:DataController"],
		beego.ControllerComments{
			Method: "GetInputFiles",
			Router: `/input/:userId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
