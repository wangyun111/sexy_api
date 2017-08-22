package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"sexy_tools/cache"
)

var mysql_url, redis_url string
var Rc *cache.Cache
var Re error
var db_name string

func init() {
	beego.Info(1111)
	runModel := beego.BConfig.RunMode
	sMap, err := beego.AppConfig.GetSection(runModel)
	if err != nil {
		panic("pattern:" + err.Error())
	}
	mysql_url = sMap["mysql_url"]
	db_name = sMap["db_name"]
	orm.RegisterDriver("mysql", orm.DRMySQL)
	err = orm.RegisterDataBase("default", "mysql", mysql_url)
	if err != nil {
		beego.Info("mysql err=" + err.Error())
	}
	redis_url = beego.AppConfig.String("beego_cache")
	Rc, Re = cache.NewCache(redis_url)
	if Re != nil {
		beego.Info("redis err=" + Re.Error())
	}
}

func init() {
	beego.Info(2222)
	o := orm.NewOrm()
	var inv []DbTable
	sql := `select table_name,table_comment from information_schema.tables  where table_schema=?`
	_, err := o.Raw(sql, db_name).QueryRows(&inv)
	if err != nil {
		beego.Info(err)
	} else {
		beego.Info(inv)
	}
	var smap map[string]interface{}
	tableSql := `show create table users`
	tableSql = `show tables`
	err = o.Raw(tableSql).QueryRow(&smap)
	beego.Info(err)
	beego.Info(smap)
}

type TableInfo struct {
	Table  string
	Table1 string
}

type DbTable struct {
	TableName    string
	TableComment string
}
