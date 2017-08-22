// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"sexy_api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/system/users", beego.NSInclude(&controllers.SysUsersController{})),
		beego.NSNamespace("/activity", beego.NSInclude(&controllers.ActivityController{})),
		beego.NSNamespace("/admin", beego.NSInclude(&controllers.AdminController{})),
		beego.NSNamespace("/cost/approve", beego.NSInclude(&controllers.CostReliefApproveController{})),
	)
	beego.AddNamespace(ns)
}
