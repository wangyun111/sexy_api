package models

import (
	// "fmt"
	"gopkg.in/mgo.v2/bson"
	"sexy_tools/time"
)

//费用减免审批
type CostReliefApprove struct {
	Id             bson.ObjectId `bson:"_id"                   json:"_id"                   description:"主键,自增"`
	Uid            int           `bson:"uid"                   json:"uid"                   description:"用户id"`
	Phone          string        `bson:"phone"                 json:"phone"                 description:"手机号"`
	State          string        `bson:"state"                 json:"state"                 description:""`
	RepSchId       int           `bson:"repSchId"              json:"repSchId"              description:"还款计划id"`
	LoanId         int           `bson:"loanId"                json:"loanId"                description:"借款id"`
	Money          float64       `bson:"money"                 json:"money"                 description:"减免金额"`
	Reason         string        `bson:"reason"                json:"reason"                description:"减免原因"`
	RequestSysId   int           `bson:"requestSysId"          json:"requestSysId"          description:"提交申请人sys_userid"`
	RequestSysName string        `bson:"requestSysName"        json:"requestSysName"        description:"提交人姓名"`
	ApproveSysId   int           `bson:"approveSysId"          json:"approveSysId"          description:"审批人sys_userid"`
	ApproveResult  string        `bson:"approveResult"         json:"approveResult"         description:"审批结果"`
	DisposeSysId   int           `bson:"disposeSysId"          json:"disposeSysId"          description:"处理人sys_userid"`
	DisposeResult  string        `bson:"disposeResult"         json:"disposeResult"         description:"处理结果"`
	CreateTime     string        `bson:"createTime"            json:"createTime"            description:"创建时间"`
	ImagesUrl      string        `bson:"imagesUrl"             json:"imagesUrl"             description:"图片详情，逗号分隔"`
}

//添加
func AddCostReliefApprove(_CostReliefApprove CostReliefApprove) (err error, id string) {
	session := GetSession()
	c := session.DB(DataBase).C("cost_relief_approve")
	bsonId := bson.NewObjectId()
	id = bsonId.Hex()
	_CostReliefApprove.Id = bsonId
	_CostReliefApprove.CreateTime = time.GetNowTime()
	err = c.Insert(&_CostReliefApprove)
	return
}

//根据id删除
func DeleteByIdCostReliefApprove(id string) (err error) {
	session := GetSession()
	c := session.DB(DataBase).C("cost_relief_approve")
	err = c.RemoveId(bson.ObjectIdHex(id))
	return
}

//条件删除
func DeleteCostReliefApprove(selectorMap map[string]interface{}) (err error) {
	session := GetSession()
	c := session.DB(DataBase).C("cost_relief_approve")
	err = c.Remove(&selectorMap)
	return
}

//根据 id 进行修改
func UpdateByIdCostReliefApprove(id bson.ObjectId, _CostReliefApprove CostReliefApprove) (err error) {
	session := GetSession()
	c := session.DB(DataBase).C("cost_relief_approve")
	err = c.UpdateId(id, &_CostReliefApprove)
	return
}

//条件修改
func UpdateCostReliefApprove(selectorMap map[string]interface{}, _CostReliefApprove CostReliefApprove) (number int, err error) {
	session := GetSession()
	c := session.DB(DataBase).C("cost_relief_approve")
	err = c.Update(&selectorMap, &_CostReliefApprove)
	return
}

//条件查询列表
func QueryCostReliefApprove(pageNo, pageSize int, sort string, selectorMap map[string]interface{}) (list []CostReliefApprove, err error) {
	session := GetSession()
	pageNo = (pageNo - 1) * pageSize
	c := session.DB(DataBase).C("cost_relief_approve")
	q := c.Find(&selectorMap).Limit(pageSize).Skip(pageNo)
	if sort != "" {
		q = q.Sort("-" + sort)
	}
	err = q.All(&list)
	return
}

//条件查询总条数
func QueryCostReliefApproveCount(selectorMap map[string]interface{}) (count int, err error) {
	session := GetSession()
	c := session.DB(DataBase).C("cost_relief_approve")
	count, err = c.Find(&selectorMap).Count()
	return
}

//条件查询单个
func FindByConditionCostReliefApprove(selectorMap map[string]interface{}) (_CostReliefApprove CostReliefApprove, err error) {
	session := GetSession()
	c := session.DB(DataBase).C("cost_relief_approve")
	err = c.Find(&selectorMap).One(&_CostReliefApprove)
	return
}

//根据 id查询单个
func FindByIdCostReliefApprove(id string) (_CostReliefApprove CostReliefApprove, err error) {
	session := GetSession()
	c := session.DB(DataBase).C("cost_relief_approve")
	err = c.FindId(bson.ObjectIdHex(id)).One(&_CostReliefApprove)
	return
}
