package models

import (
	"gopkg.in/mgo.v2/bson"
	"sexy_tools/time"
)

//会员信息
type ActivityInfo struct {
	Id         bson.ObjectId `bson:"_id"          json:"-"            description:"编号"`
	Uid        int64         `bson:"uid"          json:"uid"          description:"用户编号"`
	Title      string        `bson:"title"        json:"title"        description:"标题"`
	Cover      string        `bson:"cover"        json:"cover"        description:"封面"`
	Content    string        `bson:"content"      json:"content"      description:"内容"`
	Details    string        `bson:"details"      json:"details"      description:"详情"`
	CreateTime string        `bson:"create_time"  json:"createTime"   description:"创建时间"`
}

func AddActivity(_Activity *ActivityInfo) (has bool, msg string) {
	has = true
	msg = "success"
	session := GetSession()
	c := session.DB("wangyun").C("activity_info")
	_Activity.Id = bson.NewObjectId()
	_Activity.CreateTime = time.GetNowTime()
	err := c.Insert(_Activity)
	if err != nil {
		has = false
		msg = err.Error()
	}
	return
}

func DeleteActivity(uid int64) (has bool, msg string) {
	has = true
	msg = "success"
	var _Activity *ActivityInfo
	session := GetSession()
	c := session.DB(DataBase).C("activity_info")
	err := c.Find(&bson.M{"uid": uid}).One(&_Activity)
	if err != nil {
		has = false
		msg = err.Error()
	}
	if err == nil {
		// c.Remove(&_Activity)
		c.RemoveId(_Activity.Id)
	}
	return
}

func UpdateActivity(_Activity *ActivityInfo) (has bool, msg string) {
	has = true
	msg = "success"
	session := GetSession()
	c := session.DB(DataBase).C("activity_info")
	chageInfo, err := c.UpsertId(_Activity.Id, &_Activity)
	if err != nil {
		has = false
		msg = err.Error()
	}
	if err == nil {
		updated := chageInfo.Updated
		if updated != 1 {
			msg = "no change"
		}
	}
	return
}

func FindByIdActivity(uid int64) (has bool, msg string, _Activity ActivityInfo) {
	has = true
	msg = "success"
	session := GetSession()
	c := session.DB(DataBase).C("activity_info")
	err := c.Find(&bson.M{"uid": uid}).One(&_Activity)
	if err != nil {
		has = false
		msg = err.Error()
	}
	return
}
