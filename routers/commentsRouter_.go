package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["sexy_api/controllers:AdminController"] = append(beego.GlobalControllerRouter["sexy_api/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["sexy_api/controllers:AdminController"] = append(beego.GlobalControllerRouter["sexy_api/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["sexy_api/controllers:AdminController"] = append(beego.GlobalControllerRouter["sexy_api/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["sexy_api/controllers:AdminController"] = append(beego.GlobalControllerRouter["sexy_api/controllers:AdminController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/all`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["sexy_api/controllers:AdminController"] = append(beego.GlobalControllerRouter["sexy_api/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["sexy_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["sexy_api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["sexy_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["sexy_api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["sexy_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["sexy_api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["sexy_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["sexy_api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["sexy_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["sexy_api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
