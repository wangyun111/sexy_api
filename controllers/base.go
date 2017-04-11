package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"net/http"
	"sexy_tools/crypto"
	"sexy_tools/tools"
	"strconv"
)

var cryopt_key = "abc158sk"

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	_cid := this.Ctx.GetCookie("cid")
	if _cid == "" {
		this.Data["json"] = tools.PleaseLogin()
		this.ServeJSON()
	}
}

func (c *BaseController) ServeJSON(encoding ...bool) {
	beego.Info(encoding)
	var (
		hasIndent   = true
		hasEncoding = false
	)
	if beego.BConfig.RunMode == beego.PROD {
		hasIndent = false
	}
	if len(encoding) > 0 && encoding[0] == true {
		hasEncoding = true
	}
	c.JSON(c.Data["json"], hasIndent, hasEncoding)
}

// JSON writes json to response body.
// if coding is true, it converts utf-8 to \u0000 type.
func (c *BaseController) JSON(data interface{}, hasIndent bool, coding bool) error {
	c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	var content []byte
	var err error
	if hasIndent {
		content, err = json.MarshalIndent(data, "", "  ")
	} else {
		content, err = json.Marshal(data)
	}
	if err != nil {
		http.Error(c.Ctx.Output.Context.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return err
	}
	if coding {
		beego.Info(StringsToJSON(string(content)))
		content = []byte(StringsToJSON(string(content)))
	}
	base64Str, err := crypto.DesEncrypt(content, []byte(cryopt_key))
	desByte := crypto.Base64Encrypt(base64Str)
	return c.Ctx.Output.Body([]byte(desByte))
}

func StringsToJSON(str string) string {
	rs := []rune(str)
	jsons := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			jsons += string(r)
		} else {
			jsons += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}
	return jsons
}

// func (this *BaseController) ServeJSON(_crypto ...bool) {
// 	if len(_crypto) > 0 && _crypto[0] {
// 		this.JSON(this.Data["json"])
// 	} else {
// 	}
// }

// func (this *BaseController) JSON(data interface{}) error {
// 	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
// 	content, err := json.Marshal(data)
// 	if err != nil {
// 		return err
// 	} else {
// 		base64Str := crypto.Base64Encrypt(content)
// 		desByte, err := crypto.DesEncrypt([]byte(base64Str), []byte("12345678"))
// 		beego.Info(desByte, err)
// 		return this.Ctx.Output.Body(desByte)
// 	}
// }

func (this *BaseController) Finish() {}
