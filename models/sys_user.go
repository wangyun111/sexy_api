package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"sexy_tools/time"
)

//管理员信息
type SysUsers struct {
	Id            int64  `json:"id"                    description:"主键编号"`
	Account       string `json:"account"               description:"账号"`
	Password      string `json:"password"              description:"密码"`
	Name          string `json:"name"                  description:"真实姓名"`
	Sex           string `json:"sex"                   description:"性别"`
	Email         string `json:"email"                 description:"邮箱"`
	Phone         string `json:"phone"                 description:"电话"`
	AccountStatus string `json:"accountStatus"         description:"账号状态(inactive 禁用,active 启用)"`
	Remark        string `json:"remark"                description:"备注"`
	Code          string `json:"code"                  description:"唯一code值"`
	RoleId        int64  `json:"roleId"                description:"角色id"`
	StationId     int64  `json:"stationId"             description:"岗位id"`
	QnAccount     string `json:"qnAccount"             description:"青牛账号"`
	QnPassword    string `json:"qnPassword"            description:"青牛密码"`
	PlaceRole     string `json:"placeRole"             description:"青牛坐席角色"`
	TapeId        int64  `json:"tapeId"                description:"青牛坐席最新通话状态id"`
	AccountType   string `json:"accountType"           description:"账户类型:interior:内部,exterior:外部"`
	CreateTime    string `json:"createTime"            description:"创建时间"`
}

//添加
func AddSysUsers(_SysUsers SysUsers) (id int64, err error) {
	o := orm.NewOrm()
	_SysUsers.Id = 0
	_SysUsers.AccountStatus = "active"
	_SysUsers.AccountType = "interior"
	_SysUsers.CreateTime = time.GetNowTime()
	id, err = o.Insert(&_SysUsers)
	return
}

//根据id删除
func DeleteByIdSysUsers(id int64) (number int64, err error) {
	o := orm.NewOrm()
	sql := `delete from sys_users where id=? `
	res, err := o.Raw(sql, id).Exec()
	if err == nil {
		number, _ = res.RowsAffected()
	}
	return
}

//条件删除
func DeleteSysUsers(where string, args ...interface{}) (number int64, err error) {
	o := orm.NewOrm()
	sql := `delete from sys_users where 1=1 `
	sql += where
	res, err := o.Raw(sql, args).Exec()
	if err == nil {
		number, _ = res.RowsAffected()
	}
	return
}

//根据 id 进行修改
func UpdateByIdSysUsers(_SysUsers SysUsers) (number int64, err error) {
	o := orm.NewOrm()
	number, err = o.Update(&_SysUsers, "name", "sex", "email", "phone", "account_status",
		"remark", "code", "role_id", "station_id")
	return
}

//条件修改
func UpdateSysUsers(set, where string, args ...interface{}) (number int64, err error) {
	o := orm.NewOrm()
	sql := `update sys_users set 1=1 `
	sql += set
	sql += where
	res, err := o.Raw(sql, args).Exec()
	if err == nil {
		number, _ = res.RowsAffected()
	}
	return
}

//条件查询列表
func QuerySysUser(pageNo, pageSize int64, sort, where string, args ...interface{}) (list []SysUsers, err error) {
	o := orm.NewOrm()
	sql := `select id,account,password,name,sex,email,phone,account_status,remark,
	               code,role_id,station_id,create_time  
	        from sys_users where 1=1 `
	sql += where
	if sort != "" {
		sql += ` order by ? desc`
	}
	sql += ` limit ?,? `
	_, err = o.Raw(sql, args, (pageNo-1)*pageSize, pageSize).QueryRows(&list)
	beego.Info(err)
	return
}

//条件查询总条数
func QuerySysUserCount(where string, args ...interface{}) (count int64, err error) {
	o := orm.NewOrm()
	sql := `select count(id) from sys_users where 1=1 `
	sql += where
	err = o.Raw(sql, args).QueryRow(&count)
	return
}

//条件查询单个
func FindByConditionSysUser(where string, args ...interface{}) (_SysUsers SysUsers, err error) {
	o := orm.NewOrm()
	sql := `select id,account,password,name,sex,email,phone,account_status,remark,
	               code,role_id,station_id,create_time 
	        from sys_users where 1=1 `
	sql += where
	err = o.Raw(sql, args).QueryRow(&_SysUsers)
	return
}

//根据 id查询单个
func FindByIdSysUser(id int64) (_SysUsers SysUsers, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("sys_users")
	err = qs.Filter("id", id).One(&_SysUsers)
	return
}
