package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"sexy_tools/time"
	// "strconv"
)

//会员信息
type AdminInfo struct {
	Id         int64  `json:"-"               description:"编号"`
	Nickname   string `json:"nickName"        description:"*昵称"`
	Headimgurl string `json:"headimgurl"      description:"*头像"`
	Name       string `json:"name"            description:"姓名"`
	Phone      string `json:"phone"           description:"电话"`
	Email      string `json:"email"           description:"邮箱"`
	Address    string `json:"address"         description:"地址"`
	Account    string `json:"account"         description:"账号"`
	Password   string `json:"password"        description:"密码"`
	Remark     string `json:"remark"          description:"备注(预留信息)"`
	IsUse      string `json:"isUse"           description:"是否使用"`
	CreateTime string `json:"createTime"      description:"创建时间"`
}

func AddAdmin(_Admin *AdminInfo) (has bool, msg string) {
	has = true
	msg = "success"
	o := orm.NewOrm()
	_Admin.Id = 0
	_Admin.CreateTime = time.GetNowTime()
	_Admin.IsUse = "YES"
	_, _err := o.Insert(_Admin)
	if _err != nil {
		has = false
		msg = _err.Error()
	}
	return
}

func DeleteAdmin(id int64) (has bool, msg string) {
	has = true
	msg = "success"
	o := orm.NewOrm()
	res, _err := o.Raw("delete from sys_info where id=?", id).Exec()
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

func UpdateAdmin(_Admin *AdminInfo) (has bool, msg string) {
	has = true
	msg = "success"
	o := orm.NewOrm()
	sql := "update admin_info set id=id, "
	id := _Admin.Id
	nickname := _Admin.Nickname
	headimgurl := _Admin.Headimgurl
	name := _Admin.Name
	address := _Admin.Address
	isUse := _Admin.IsUse
	if nickname != "" {
		sql += " ,nickname='" + nickname + "'"
	}
	if headimgurl != "" {
		sql += " ,headimgurl='" + headimgurl + "'"
	}
	if name != "" {
		sql += " ,name='" + name + "'"
	}
	if address != "" {
		sql += " ,address='" + address + "'"
	}
	if isUse != "" {
		sql += " ,is_use='" + isUse + "'"
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

// func QueryAdmin(pageNo, pageSize int64, nickname, name, phone string) (has bool, msg string, page *Page) {
// 	has = true
// 	msg = "success"
// 	if pageNo < 1 {
// 		pageNo = 1
// 	}
// 	if pageSize < 1 {
// 		pageSize = 20
// 	}
// 	o := orm.NewOrm()
// 	var _err error
// 	var countModel CountModel
// 	countSql := " select count(*) as count from admin_info where 1=1 "
// 	sql := "select id,nickname,headimgurl,name,phone,email,address,account,password,remark,is_use,create_time from admin_info where 1=1 "
// 	if nickname != "" {
// 		countSql += " and nickname like '%" + nickname + "%'"
// 		sql += " and nickname like '%" + nickname + "%'"
// 	}
// 	if name != "" {
// 		countSql += " and name  like '%" + name + "%'"
// 		sql += " and name like '%" + name + "%'"
// 	}
// 	if phone != "" {
// 		countSql += " and phone like '%" + phone + "%'"
// 		sql += " and phone like '%" + phone + "%'"
// 	}
// 	_err = o.Raw(countSql).QueryRow(&countModel)
// 	beego.Info(_err)
// 	count := countModel.Count
// 	if count == 0 {
// 		has = false
// 		msg = ""
// 		if _err != nil {
// 			msg = _err.Error()
// 		}
// 		return
// 	}
// 	sql += " order by create_time desc LIMIT " + strconv.FormatInt(pageSize, 10) + " OFFSET " + strconv.FormatInt((pageNo-1)*pageSize, 10)
// 	var list []AdminInfo
// 	_, _err = o.Raw(sql).QueryRows(&list)
// 	beego.Info(_err)
// 	if _err != nil {
// 		has = false
// 		msg = _err.Error()
// 		return
// 	}
// 	page = PageUtil(count, pageNo, pageSize, list)
// 	return
// }

func FindByIdAdmin(id int64) (has bool, msg string, _Admin AdminInfo) {
	has = true
	msg = "success"
	o := orm.NewOrm()
	sql := "select id,nickname,headimgurl,name,phone,email,address,account,password,remark,is_use,create_time from admin_info where id=? "
	err := o.Raw(sql, id).QueryRow(&_Admin)
	beego.Info(err)
	if err != nil {
		has = false
		msg = ""
		if err != orm.ErrNoRows {
			msg = err.Error()
		}
	}
	return
}

func LoginAdmin(account, password string) (has bool, msg string, _Admin AdminInfo) {
	has = true
	msg = "success"
	o := orm.NewOrm()
	sql := "select * from admin_info where account=? and password=? "
	err := o.Raw(sql, account, password).QueryRow(&_Admin)
	beego.Info(err)
	if err != nil {
		has = false
		msg = ""
		if err != orm.ErrNoRows {
			msg = err.Error()
		}
	}
	return
}
