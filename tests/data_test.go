package test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	daoApi "api/dao_service"
	_ "app-service/data-service/routers"
	"model"
)

const (
	base_url = "http://localhost:8080/v1/data"
)

func Test_Data_GetInputFiles(t *testing.T) {
	// create admin
	daoApi.UserDaoApi.Init("http://user-dao-service:8080")
	var admin model.User
	admin.Id = 0
	admin.Name = "admin"
	newAdmin, err := daoApi.UserDaoApi.Create(&admin)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log("admin:", *newAdmin)
	// create data input directory
	path := "/pme2017/workspace/" + admin.Name + "/data/input"
	err = os.MkdirAll(path, os.ModePerm) //生成多级目录

	// create input data
	fn := path + "/test_data.sgy"
	f, err := os.Create(fn) //创建文件
	defer f.Close()
	if err != nil {
		t.Log(err)
		return
	}

	// get input file
	res, err := http.Get(base_url + "/1")
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))

	var response model.Response
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if response.Reason == "success" {
		t.Log("PASS OK")
	} else {
		t.Log("ERROR:", response.Reason)
		t.FailNow()
	}
}
