package service

import (
	"model"

	"fmt"
	"io/ioutil"
	"os"

	daoApi "api/dao_service"

	"github.com/astaxie/beego"
)

type DataService struct {
}

func (this *DataService) GetInputFiles(userId int64) ([]string, error) {
	var err error
	var inputfiles []string

	// get user
	beego.Debug("->get user")
	var user *model.User
	user, err = daoApi.UserDaoApi.GetById(userId)
	if err != nil {
		beego.Debug("get user failed")
		err = fmt.Errorf("%s", "get user failed")
		return nil, err
	}

	// 查看目录下文件内容
	beego.Debug("->read dir")
	var dir []os.FileInfo
	var cfg = beego.AppConfig
	dirPath := cfg.String("workspace") + "/" + user.Name + "/data/input"

	dir, err = ioutil.ReadDir(dirPath)
	if err != nil {
		beego.Debug("read dir failed")
		return nil, err
	}

	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}

		inputfiles = append(inputfiles, fi.Name())
	}

	beego.Debug("result:", inputfiles)

	return inputfiles, err
}
