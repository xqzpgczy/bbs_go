package api

import (
	"bbs-go/model/constants"
	"bbs-go/pkg/config"
	"fmt"
	"github.com/kataras/iris/v12"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/mlogclub/simple/web"
	"github.com/sirupsen/logrus"

	"bbs-go/services"
)

type UploadController struct {
	Ctx iris.Context
}

func (c *UploadController) Post() *web.JsonResult {
	user := services.UserTokenService.GetCurrent(c.Ctx)
	if err := services.UserService.CheckPostStatus(user); err != nil {
		return web.JsonError(err)
	}

	file, header, err := c.Ctx.FormFile("image")
	if err != nil {
		return web.JsonErrorMsg(err.Error())
	}
	defer file.Close()

	if header.Size > constants.UploadMaxBytes {
		return web.JsonErrorMsg("图片不能超过" + strconv.Itoa(constants.UploadMaxM) + "M")
	}

	//contentType := header.Header.Get("Content-Type")
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return web.JsonErrorMsg(err.Error())
	}

	image_path := "." + config.Instance.StaticPath + "/" + "www/upload/image/"
	image_name := strconv.Itoa(time.Now().Nanosecond()) + "_" + header.Filename

	err = ioutil.WriteFile(image_path+image_name, fileBytes, 0777)

	logrus.Info("上传文件：", header.Filename, " size:", header.Size)
	fmt.Println(err)

	//url, err := uploader.PutImage(fileBytes, contentType)
	//if err != nil {
	//	return web.JsonErrorMsg(err.Error())
	//}

	url := config.Instance.BaseUrl + "/api/www/upload/image/" + image_name
	return web.NewEmptyRspBuilder().Put("url", url).JsonResult()
}
