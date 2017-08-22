package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"sexy_api/models"
	"sexy_tools/tools"
)

// 管理员模块
type SysUsersController struct {
	beego.Controller
}

// @Summary 添加管理员
// @Description
// @Param	body    body 	      models.SysUsers	true		"管理员信息"
// @Success 200     {string}      code              200
// @Failure 201     {string}      code              201
// @router / [post]
func (this *SysUsersController) Post() {
	errmsg := ""
	defer func() {
		if errmsg != "" {
			errmsg = "SysUsersController,Post:" + errmsg
			beego.Info(errmsg)
		}
		this.ServeJSON()
	}()
	var _SysUsers models.SysUsers
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &_SysUsers)
	if err != nil {
		this.Data["json"] = tools.ParamErrorMsg(err.Error())
		return
	}
	account := _SysUsers.Account
	password := _SysUsers.Password
	if account == "" || password == "" {
		this.Data["json"] = tools.ParamNull()
		return
	}
	_, err = models.AddSysUsers(_SysUsers)
	if err != nil {
		errmsg += err.Error()
		this.Data["json"] = tools.OperationFalseMsg(err.Error())
		return
	}
	this.Data["json"] = tools.OperationSuccess()
}

// @Summary 删除管理员
// @Description
// @Param	id     query 	  int	    true		"主键 id"
// @Success 200    {string}   code      200
// @Failure 201    {string}   code      201
// @router / [delete]
func (this *SysUsersController) Delete() {
	errmsg := ""
	defer func() {
		if errmsg != "" {
			errmsg = "SysUsersController,Delete:" + errmsg
			beego.Info(errmsg)
		}
		this.ServeJSON()
	}()
	id, _ := this.GetInt64("id")
	if id < 1 {
		this.Data["json"] = tools.ParamNull()
		return
	}
	_, err := models.DeleteByIdSysUsers(id)
	if err != nil {
		errmsg += err.Error()
		this.Data["json"] = tools.OperationFalseMsg(err.Error())
		return
	}
	this.Data["json"] = tools.OperationSuccess()
}

// @Summary 修改管理员信息
// @Description
// @Param	body   body 	models.SysUsers	  true		"管理员信息"
// @Success 200 {string} code 200
// @Failure 201 {string} code 201
// @router / [put]
func (this *SysUsersController) Put() {
	errmsg := ""
	defer func() {
		if errmsg != "" {
			errmsg = "SysUsersController,Put:" + errmsg
			beego.Info(errmsg)
		}
		this.ServeJSON()
	}()
	var _SysUsers models.SysUsers
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &_SysUsers)
	if err != nil {
		this.Data["json"] = tools.ParamErrorMsg(err.Error())
		return
	}
	id := _SysUsers.Id
	if id < 1 {
		this.Data["json"] = tools.ParamNullMsg("id不能为空")
		return
	}
	_, err = models.UpdateByIdSysUsers(_SysUsers)
	if err != nil {
		errmsg += err.Error()
		this.Data["json"] = tools.ParamErrorMsg(err.Error())
		return
	}
	this.Data["json"] = tools.OperationSuccess()
}

// @Summary 获取管理员列表信息
// @Description
// @Param	pageNo		query    	int 	      false		"页数"
// @Param	pageSize	query    	int 	      flase		"条数"
// @Param	name	    query 	    string	      false		"姓名"
// @Param	account	    query 	    string	      false		"账号"
// @Param	sort	    query 	    string	      false		"排序方式(传了createTime 时间倒序)"
// @Success 200         {object}    models.SysUsers
// @Failure 203         {string}    code 203
// @router /all [get]
func (this *SysUsersController) GetAll() {
	errmsg := ""
	defer func() {
		if errmsg != "" {
			errmsg = "SysUsersController,GetAll:" + errmsg
			beego.Info(errmsg)
		}
		this.ServeJSON()
	}()
	args := []interface{}{}
	where := ""
	name := this.GetString("name")
	if name != "" {
		args = append(args, "'%"+name+"%'")
		where += " and name like ?"
	}
	account := this.GetString("account")
	if account != "" {
		args = append(args, "'%"+account+"%'")
		where += " and account like ?"
	}
	count, err := models.QuerySysUserCount(where, args...)
	if err != nil {
		errmsg += err.Error()
	}
	if count == 0 {
		this.Data["json"] = tools.ReturnDataNull()
		return
	}
	sort := this.GetString("sort")
	if sort != "" {
		args = append(args, sort)
	}
	pageNo, _ := this.GetInt64("pageNo")
	if pageNo < 1 {
		pageNo = 1
	}
	pageSize, _ := this.GetInt64("pageSize")
	if pageSize < 1 {
		pageSize = 20
	}
	list, err := models.QuerySysUser(pageNo, pageSize, sort, where, args...)
	if err != nil {
		errmsg += err.Error()
	}
	page := models.PageUtil(count, pageNo, pageSize, list)
	this.Data["json"] = page
}

// @Summary 获取单个管理员信息
// @Description
// @Param	 id	     query 	    int	      true		"主键 id"
// @Success 200      {object} models.SysUsers
// @Failure 203      {string} code 203
// @router / [get]
func (this *SysUsersController) Get() {
	errmsg := ""
	defer func() {
		if errmsg != "" {
			errmsg = "SysUsersController,Get:" + errmsg
			beego.Info(errmsg)
		}
		this.ServeJSON()
	}()
	id, _ := this.GetInt64("id")
	if id < 1 {
		this.Data["json"] = tools.ParamNull()
		return
	}
	_SysUsers, err := models.FindByIdSysUser(id)
	if err != nil {
		errmsg = err.Error()
		if errmsg != "<QuerySeter> no row found" {
			errmsg += err.Error()
		}
		this.Data["json"] = tools.ReturnDataNull()
		return
	}
	this.Data["json"] = _SysUsers
}
