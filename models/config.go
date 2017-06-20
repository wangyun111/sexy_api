package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"sexy_tools/cache"
)

var mysql_url, redis_url string
var Rc *cache.Cache
var Re error

func init() {
	runModel := beego.BConfig.RunMode
	sMap, err := beego.AppConfig.GetSection(runModel)
	if err != nil {
		panic("pattern:" + err.Error())
	}
	mysql_url = sMap["mysql_url"]
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", mysql_url)
	// redis_url = beego.AppConfig.String("beego_cache")
	// beego.Info(redis_url)
	// Rc, Re = cache.NewCache(redis_url)
	// beego.Info(Re)
}
