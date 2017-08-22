package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"sexy_api/models"
	"sexy_tools/tools"
)

// 费用减免模块
type CostReliefApproveController struct {
	beego.Controller
}

// @Summary 添加费用减免信息
// @Description
// @Param	body   body 	models.CostReliefApprove	true		"费用减免信息信息"
// @Success 200 {string} code 200
// @Failure 201 {string} code 201
// @router / [post]
func (this *CostReliefApproveController) Post() {
	var msg string
	defer this.ServeJSON()
	var _CostReliefApprove models.CostReliefApprove
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &_CostReliefApprove)
	if err != nil {
		msg += err.Error()
		this.Data["json"] = tools.ParamErrorMsg(msg)
		return
	}
	err, _ = models.AddCostReliefApprove(_CostReliefApprove)
	if err != nil {
		msg += err.Error()
		this.Data["json"] = tools.OperationFalseMsg(msg)
		return
	}
	this.Data["json"] = tools.OperationSuccess()
}

// @Summary 删除费用减免信息
// @Description
// @Param	id	  query 	 string	    true		"主键 id"
// @Success 200   {string}   code   200
// @Failure 201   {string}   code   201
// @router / [delete]
func (this *CostReliefApproveController) Delete() {
	var msg string
	defer this.ServeJSON()
	id := this.GetString("id")
	if id == "" || len(id) != 24 {
		this.Data["json"] = tools.ParamNull()
		return
	}
	err := models.DeleteByIdCostReliefApprove(id)
	if err != nil {
		msg += err.Error()
		this.Data["json"] = tools.OperationFalseMsg(msg)
		return
	}
	this.Data["json"] = tools.OperationSuccess()
}

// @Summary 修改费用减免信息
// @Description
// @Param	body   body 	models.CostReliefApprove	true		"费用减免信息信息"
// @Success 200 {string} code 200
// @Failure 201 {string} code 201
// @router / [put]
func (this *CostReliefApproveController) Put() {
	var msg string
	defer this.ServeJSON()
	var _CostReliefApprove models.CostReliefApprove
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &_CostReliefApprove)
	if err != nil {
		msg += err.Error()
		this.Data["json"] = tools.ParamErrorMsg(msg)
		return
	}
	id := _CostReliefApprove.Id
	idHex := id.Hex()
	beego.Info(id == "")
	beego.Info(idHex)
	if len(idHex) != 24 {
		this.Data["json"] = tools.ParamNull()
		return
	}
	err = models.UpdateByIdCostReliefApprove(id, _CostReliefApprove)
	if err != nil {
		msg += err.Error()
		this.Data["json"] = tools.OperationFalseMsg(msg)
		return
	}
	this.Data["json"] = tools.OperationSuccess()
}

// @Summary 获取费用减免信息
// @Description
// @Param	pageNo		query    	int 	      false		"页数"
// @Param	pageSize	query    	int 	      flase		"条数"
// @Param	state	    query 	    string	      false		"状态"
// @Param	sort	    query 	    string	      false		"排序方式"
// @Success 200 {object}   models.CostReliefApprove
// @Failure 204 {string}   code 204
// @router /all [get]
func (this *CostReliefApproveController) GetAll() {
	var msg string
	defer this.ServeJSON()
	selectorMap := map[string]interface{}{}
	pageNo, _ := this.GetInt("pageNo")
	pageSize, _ := this.GetInt("pageSize")
	state := this.GetString("state")
	sort := this.GetString("sort")
	if pageNo < 1 {
		pageNo = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	if state != "" {
		selectorMap["state"] = state
	}
	count, _ := models.QueryCostReliefApproveCount(selectorMap)
	if count == 0 {
		this.Data["json"] = tools.ReturnDataNull()
		return
	}
	list, err := models.QueryCostReliefApprove(pageNo, pageSize, sort, selectorMap)
	if err != nil {
		msg += err.Error()
		this.Data["json"] = tools.RoutineError(msg)
		return
	}
	page := models.PageUtil(int64(count), int64(pageNo), int64(pageSize), list)
	this.Data["json"] = page
}

// @Summary 查询单个费用减免信息
// @Description
// @Param	id	query 	string	  true	 "id"
// @Success 200 {object}   models.CostReliefApprove
// @Failure 204 {string}   code 204
// @router / [get]
func (this *CostReliefApproveController) Get() {
	var msg string
	defer this.ServeJSON()
	id := this.GetString("id")
	if id == "" || len(id) != 24 {
		this.Data["json"] = tools.ParamNull()
		return
	}
	_CostReliefApprove, err := models.FindByIdCostReliefApprove(id)
	if err != nil {
		msg += err.Error()
		this.Data["json"] = tools.GetReturnObject(msg)
	}
	this.Data["json"] = _CostReliefApprove
}
