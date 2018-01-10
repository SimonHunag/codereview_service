package controllers

import (
	"github.com/astaxie/beego"
	"codereview/models"
	"encoding/json"
	"strconv"
)

// Operations about Project
type ProjectController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all Project
// @Success 200 {object} models.Project
// @router / [get]
func (u *ProjectController) GetAll() {
	project := models.GetAllProject()
	u.Data["json"] = project
	u.ServeJSON()
}

// @Title CreateProject
// @Description create project
// @Param	body		body 	models.Project	true		"body for project content"
// @Success 200 {int} models.Project.ProjectId
// @Failure 403 body is empty
// @router / [post]
func (c *ProjectController) Post() {
	var project models.Project
	json.Unmarshal(c.Ctx.Input.RequestBody, &project)

	newproject, err := models.NewProject(&project)

	if err != nil {
		beego.Error("NewProject:", err)
		c.Data["json"] = models.NewErrorInfo(ErrSystem)
		c.ServeJSON()
		return
	}

	var Result=models.CodeInfo{}
	if code, err := newproject.Insert(); err != nil {
		beego.Error("InsertProject:", err)
		if code == models.ErrDupRows {
			Result = models.CodeInfo{
				Code: 0,
				Info: ErrDupProject,
			}
		} else {
			Result = models.CodeInfo{
				Code: 0,
				Info: ErrDatabase,
			}
		}
	} else {
		 Result = models.CodeInfo{
			Code: 1,
			Info: strconv.FormatInt(code, 10),
		}
	}

	c.Data["json"] = Result
	c.ServeJSON()
	return

}
