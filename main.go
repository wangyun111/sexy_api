package main

import (
	"github.com/astaxie/beego"
	_ "sexy_api/routers"
	_ "sexy_api/utils"
)

func main() {
	// beego.BeeLogger.DelLogger("console")
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 600
	beego.Run()
}
