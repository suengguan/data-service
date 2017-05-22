package controllers

import (
	"app-service/data-service/models"
	"app-service/data-service/service"
	"encoding/json"
	"fmt"
	"model"

	"github.com/astaxie/beego"
)

// Operations about Data
type DataController struct {
	beego.Controller
}

// @Title GetInputFiles
// @Description get input files
// @Param	userid		path 	int64	true		"The key for staticblock"
// @Success 200 {object} models.Response
// @router /input/:userId [get]
func (this *DataController) GetInputFiles() {
	var err error
	var response models.Response

	var userId int64
	userId, err = this.GetInt64(":userId")
	//beego.Debug("GetInputFiles", userId)
	if userId > 0 && err == nil {
		var svc service.DataService
		var inputfiles []string
		var result []byte
		inputfiles, err = svc.GetInputFiles(userId)
		if err == nil {
			result, err = json.Marshal(&inputfiles)
			if err == nil {
				response.Status = model.MSG_RESULTCODE_SUCCESS
				response.Reason = "success"
				response.Result = string(result)
			}
		}
	} else {
		beego.Debug(err)
		err = fmt.Errorf("%s", "user id is invalid")
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}
	this.Data["json"] = &response

	this.ServeJSON()
}
