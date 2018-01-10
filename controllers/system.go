package controllers

import (
	"github.com/astaxie/beego"
	"codereview/models"
	"encoding/json"
)

// Operations about System
type SystemController struct {
	beego.Controller
}

// @Title Get
// @Description get system by code
// @Param	code		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.System
// @Failure 403 :code is empty
// @router /:code [get]
func (u *SystemController) Get() {
	code := u.GetString(":code")
	if code != "" {
		system, err := models.GetSystem(code)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = system
		}
	}
	u.ServeJSON()
}

// @Title CreateSystem
// @Description create system
// @Param	body		body 	models.SystemForm	true		"body for project content"
// @Success 200 {object} models.CodeInfo
// @Failure 403 body is empty
// @router / [post]
func (c *SystemController) Post() {
	var systemForm models.SystemForm
	var Result = models.CodeInfo{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &systemForm)

	system, err := models.NewSystem(&systemForm);
	if err != nil {
		Result = models.CodeInfo{
			Code: 0,
			Info: ErrInputData,
		}
		c.Data["json"] = Result
		c.ServeJSON()
		return
	}

	if err := system.Insert(); err != nil {
		beego.Error("InsertSystem:", err)
		Result = models.CodeInfo{
			Code: 0,
			Info: err.Error(),
		}

	} else {
		Result = models.CodeInfo{
			Code: 1,
			Info: "success",
		}
	}

	c.Data["json"] = Result
	c.ServeJSON()
	return

}
