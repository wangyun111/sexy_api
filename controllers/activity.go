package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"sexy_api/models"
	"sexy_api/utils"
	"sexy_tools/tools"
	"strconv"
	"time"
)

// 会员模块
type ActivityController struct {
	beego.Controller
}

// @Summary 添加活动
// @Description
// @Param	body   body 	models.ActivityInfo	true		"活动信息"
// @Success 200 {string} code 200
// @Failure 201 {string} code 201
// @router / [post]
func (this *ActivityController) Post() {
	defer this.ServeJSON()
	var _Activity models.ActivityInfo
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &_Activity)
	if err != nil {
		this.Data["json"] = tools.ParamErrorMsg(err.Error())
	} else {
		title := _Activity.Title
		content := _Activity.Content
		if title == "" || content == "" {
			this.Data["json"] = tools.ParamNull()
			return
		}
		has, msg := models.AddActivity(&_Activity)
		if has {
			this.Data["json"] = tools.OperationSuccess()
		}
		if !has {
			this.Data["json"] = tools.OperationFalseMsg(msg)
		}
	}
}

// @Summary 删除活动
// @Description
// @Param	uid	     query 	int	    true		"uid"
// @Success 200   {string}   code   200
// @Failure 201   {string}   code   201
// @router / [delete]
func (this *ActivityController) Delete() {
	defer this.ServeJSON()
	uid, _ := this.GetInt64("uid")
	if uid < 1 {
		this.Data["json"] = tools.ParamNull()
	} else {
		has, msg := models.DeleteActivity(uid)
		if has {
			this.Data["json"] = tools.OperationSuccess()
		} else {
			this.Data["json"] = tools.OperationFalseMsg(msg)
		}
	}
}

// @Summary 修改活动
// @Description
// @Param	body   body 	models.ActivityInfo	true		"活动信息"
// @Success 200 {string} code 200
// @Failure 201 {string} code 201
// @router / [put]
func (this *ActivityController) Put() {
	defer this.ServeJSON()
	var _Activity models.ActivityInfo
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &_Activity)
	if err != nil {
		this.Data["json"] = tools.ParamErrorMsg(err.Error())
	} else {
		if id := _Activity.Id; id == "" {
			this.Data["json"] = tools.ParamNull()
		} else {
			has, msg := models.UpdateActivity(&_Activity)
			if has {
				this.Data["json"] = tools.OperationSuccess()
			} else {
				this.Data["json"] = tools.OperationFalseMsg(msg)
			}
		}
	}
}

// @Summary 获取活动信息
// @Description
// @Param	uid	     query 	int	    true		"uid"
// @Success 200 {object}   models.ActivityInfo
// @Failure 204 {string}   code 204
// @router / [get]
func (this *ActivityController) Get() {
	defer this.ServeJSON()
	uid, _ := this.GetInt64("uid")
	if uid < 1 {
		this.Data["json"] = tools.ParamNull()
	}
	if uid > 0 {
		status := utils.Rc.Lock(strconv.FormatInt(uid, 10), 3*time.Second)
		beego.Info(status)
		if status {
			has, msg, _Activity := models.FindByIdActivity(uid)
			if has {
				this.Data["json"] = _Activity
			}
			if !has {
				this.Data["json"] = tools.GetReturnObject(msg)
			}
		} else {
			this.Data["json"] = tools.OperationFalseMsg("请勿频繁操作")
		}

	}
}
