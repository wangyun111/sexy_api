package logs

import (
	"github.com/astaxie/beego/logs"
)

var BeegoLog *logs.BeeLogger

func init() {
	//logs.EnableFuncCallDepth(true)
	// logs.SetLogger("console")
	//logs.SetLogger(logs.AdapterFile,
	//	`{"filename":"./logs/sexy_api.log","level":4,"maxlines":10000,"maxsize":33554432,"daily":true,"maxdays":10}`)
	//BeegoLog = logs.GetBeeLogger()
}
