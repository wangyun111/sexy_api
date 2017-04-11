package main

import (
	"github.com/astaxie/beego"
	_ "sexy_api/logs"
	_ "sexy_api/routers"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 600
	beego.Run()
}
