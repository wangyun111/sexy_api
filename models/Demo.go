package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"sexy_tools/time"
	"strconv"
)

const (
	sys_demo string = "sys_demo"
)

//会员信息
type Demo struct {
	Id         int64   `json:"-"                         description:"编号"`
	Uid        int64   `json:"uid"                       description:"uid"`
	Score      float64 `json:"score"                     description:"分数"`
	Name       string  `json:"name"                      description:"姓名"`
	CreateTime string  `json:"createTime"                description:"创建时间"`
}

func AddDemo(_Demo *Demo) (has bool, msg string) {
	has = true
	msg = "success"
	o := orm.NewOrm()
	_Demo.Id = 0
	_Demo.CreateTime = time.GetNowTime()
	_, _err := o.Insert(_Demo)
	if _err != nil {
		has = false
		msg = _err.Error()
	}
	return
}

func DeleteDemo(id int64) (has bool, msg string) {
	has = true
	msg = "success"
	o := orm.NewOrm()
	res, _err := o.Raw("delete from "+sys_demo+" where id=?", id).Exec()
	if _err != nil {
		has = false
		msg = _err.Error()
	} else {
		num, _ := res.RowsAffected()
		if num == 0 {
			msg = ""
		}
	}
	return
}

func UpdateDemo(_Demo *Demo) (has bool, msg string) {
	has = true
	msg = "success"
	o := orm.NewOrm()
	sql := "update " + sys_demo + " set 1=1, "
	id := _Demo.Id
	name := _Demo.Name
	score := _Demo.Score
	uid := _Demo.Uid
	if name != "" {
		sql += " ,name='" + sys_demo + "'"
	}
	if score > 0 {
		sql += " ,score=" + strconv.FormatFloat(score, 'f', 2, 64)
	}
	if uid > 0 {
		sql += " ,uid=" + strconv.FormatInt(uid, 10)
	}
	sql += " where id=?"
	res, _err := o.Raw(sql, id).Exec()
	if _err != nil {
		has = false
		msg = _err.Error()
	} else {
		num, _ := res.RowsAffected()
		if num == 0 {
			msg = ""
		}
	}
	return
}

func QueryDemo(pageNo, pageSize int64, uid int64, name string) (has bool, msg string, page *Page) {
	has = true
	msg = "success"
	if pageNo < 1 {
		pageNo = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	o := orm.NewOrm()
	var _err error
	var countModel CountModel
	countSql := " select count(*) as count from " + sys_demo + " where 1=1 "
	sql := "select * from " + sys_demo + " where 1=1 "
	if uid > 0 {
		countSql += " and uid =" + strconv.FormatInt(uid, 10)
		sql += " and uid =" + strconv.FormatInt(uid, 10)
	}
	if name != "" {
		countSql += " and name  like '%" + name + "%'"
		sql += " and name like '%" + name + "%'"
	}
	_err = o.Raw(countSql).QueryRow(&countModel)
	beego.Info(_err)
	count := countModel.Count
	if count == 0 {
		has = false
		msg = ""
		if _err != nil {
			msg = _err.Error()
		}
		return
	}
	sql += " order by create_time desc LIMIT " + strconv.FormatInt(pageSize, 10) + " OFFSET " + strconv.FormatInt((pageNo-1)*pageSize, 10)
	var list []Demo
	_, _err = o.Raw(sql).QueryRows(&list)
	beego.Info(_err)
	if _err != nil {
		has = false
		msg = _err.Error()
		return
	}
	page = PageUtil(count, pageNo, pageSize, list)
	return
}

func FindByIdDemo(id int64) (has bool, msg string, _Demo Demo) {
	has = true
	msg = "success"
	o := orm.NewOrm()
	sql := "select * from " + sys_demo + " where id=? "
	err := o.Raw(sql, id).QueryRow(&_Demo)
	if err != nil {
		has = false
		msg = ""
		if err != orm.ErrNoRows {
			msg = err.Error()
		}
	}
	return
}
