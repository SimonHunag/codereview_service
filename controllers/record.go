package controllers

import (
	"github.com/astaxie/beego"
	"codereview/models"
	"encoding/json"
)

// Operations about Record
type RecordController struct {
	beego.Controller
}

// @Title CreateRecord
// @Description create record
// @Param	body		body 	models.RecordForm	true		"body for record content"
// @Success 200 {int} models.RecordListResult
// @Failure 403 body is empty
// @router / [post]
func (u *RecordController) Post() {
	var rform models.RecordForm
	json.Unmarshal(u.Ctx.Input.RequestBody, &rform)
	record := models.NewRecord(&rform)
	err := record.AddRecord()
	if err != nil {
		result := models.NewErrorInfo(err.Error())
		u.Data["json"] = result
		u.ServeJSON()
		return
	}

	records, err := models.GetAllRecordByCode(record.SystemCode)
	if err != nil {
		result := models.NewErrorInfo(err.Error())
		u.Data["json"] = result
		u.ServeJSON()
		return
	}
	result := models.RecordListResult{
		Records: records,
		CodeInfo: models.CodeInfo{
			Code: 1,
			Info: "success",
		},
	}
	u.Data["json"] = result
	u.ServeJSON()
}

// @Title Get
// @Description get system by code
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.RecordResult
// @Failure 403 :id is empty
// @router /:id [get]
func (u *RecordController) Get() {
	id, err := u.GetInt64(":id")
	if err == nil {
		record, err := models.GetRecordByID(id)
		if err != nil {
			u.Data["json"] = models.NewErrorInfo(err.Error())
		} else {
			u.Data["json"] = models.NewRecordResult(record)
		}
	} else {
		result := models.NewErrorInfo(ErrInputData)
		u.Data["json"] = result
		u.ServeJSON()
	}
	u.ServeJSON()
}
