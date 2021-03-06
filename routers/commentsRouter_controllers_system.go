package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["sexy_api/controllers/system:AdminController"] = append(beego.GlobalControllerRouter["sexy_api/controllers/system:AdminController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["sexy_api/controllers/system:AdminController"] = append(beego.GlobalControllerRouter["sexy_api/controllers/system:AdminController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["sexy_api/controllers/system:AdminController"] = append(beego.GlobalControllerRouter["sexy_api/controllers/system:AdminController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["sexy_api/controllers/system:AdminController"] = append(beego.GlobalControllerRouter["sexy_api/controllers/system:AdminController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
