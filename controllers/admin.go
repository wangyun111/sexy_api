package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"sexy_api/models"
	"sexy_tools/tools"
)

// 会员模块
type AdminController struct {
	BaseController
}

// @Summary 添加会员
// @Description
// @Param	body		body 	models.AdminInfo	true		"会员信息"
// @Success 200 {string} code 200
// @Failure 201 {string} code 201
// @router / [post]
func (this *AdminController) Post() {
	defer this.ServeJSON()
	var admin models.AdminInfo
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &admin)
	if err != nil {
		this.Data["json"] = tools.ParamError(err.Error())
	} else {
		nickname := admin.Nickname
		headimgurl := admin.Headimgurl
		if nickname == "" || headimgurl == "" {
			this.Data["json"] = tools.ParamNull()
		} else {
			has, msg := models.AddAdmin(&admin)
			if has {
				this.Data["json"] = tools.OperationSuccess()
			} else {
				this.Data["json"] = tools.OperationFalseMsg(msg)
			}
		}
	}
}

// @Summary 删除会员
// @Description
// @Param	id		  query 	int  	true		"会员编号"
// @Success 200 {string} code 200
// @Failure 201 {string} code 201
// @router / [delete]
func (this *AdminController) Delete() {
	defer this.ServeJSON()
	id, _ := this.GetInt64("id")
	if id < 1 {
		this.Data["json"] = tools.ParamNull()
	} else {
		has, msg := models.DeleteAdmin(id)
		if has {
			this.Data["json"] = tools.OperationSuccess()
		} else {
			this.Data["json"] = tools.OperationFalseMsg(msg)
		}
	}
}

// @Summary 修改会员
// @Description
// @Param	body		body 	models.AdminInfo	true		"会员信息"
// @Success 200 {string} code 200
// @Failure 201 {string} code 201
// @router / [put]
func (this *AdminController) Put() {
	defer this.ServeJSON()
	var admin models.AdminInfo
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &admin)
	if err != nil {
		this.Data["json"] = tools.ParamError(err.Error())
	} else {
		if id := admin.Id; id < 1 {
			this.Data["json"] = tools.ParamNull()
		} else {
			has, msg := models.UpdateAdmin(&admin)
			if has {
				this.Data["json"] = tools.OperationSuccess()
			} else {
				this.Data["json"] = tools.OperationFalseMsg(msg)
			}
		}
	}
}

// @Summary 分页查询会员信息
// @Description
// @Param	pageNo		query    	int 	false		"页数"
// @Param	pageSize	query    	int 	flase		"条数"
// @Param	nickname    query    	string 	flase		"昵称"
// @Param	name	    query    	string 	flase		"姓名"
// @Param	phone    	query 	    string	false		"电话"
// @Success 200 {object} models.Page
// @Failure 204 {string} code 204
// @router /all [get]
func (this *AdminController) GetAll() {
	defer this.ServeJSON()
	pageNo, _ := this.GetInt64("pageNo")
	pageSize, _ := this.GetInt64("pageSize")
	nickname := this.GetString("nickname")
	name := this.GetString("name")
	phone := this.GetString("phone")
	has, msg, page := models.QueryAdmin(pageNo, pageSize, nickname, name, phone)
	if has {
		this.Data["json"] = page
	} else {
		if msg != "" {
			this.Data["json"] = tools.OperationFalseMsg(msg)
		} else {
			this.Data["json"] = tools.ReturnDataNull()
		}
	}
}

// @Summary 获取会员信息
// @Description
// @Param	id	       query 	int	    true		"会员编号"
// @Success 200 {object}   models.AdminInfo
// @Failure 204 {string}   code 204
// @router / [get]
func (this *AdminController) Get() {
	// this.Ctx.Output.Body([]byte("1111"))
	beego.Info(models.Rc.Get("a"))
	// this.Ctx.WriteString)
	// this.Ctx.SetCookie("cookie", "1111")
	// this.SetSession("uid", 112)
	defer this.ServeJSON()
	id, _ := this.GetInt64("id")
	if id < 1 {
		this.Data["json"] = tools.ParamNull()
	} else {
		has, msg, admin := models.FindByIdAdmin(id)
		if has {
			this.Data["json"] = admin
		} else {
			if msg != "" {
				this.Data["json"] = tools.OperationFalseMsg(msg)
			} else {
				this.Data["json"] = tools.ReturnDataNull()
			}

		}
	}
}
