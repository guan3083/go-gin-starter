package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-starter/models"
	"go-gin-starter/pkg/logf"
	"go-gin-starter/pkg/setting"
	"go-gin-starter/pkg/util"
	vali "go-gin-starter/pkg/validation"
	"go-gin-starter/routers"
	"log"
	"net/http"
)

func init() {
	setting.Setup()
	logf.Setup()
	models.Setup()
	util.Setup()
	vali.InitValidation()
}

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/EDDYCJY/go-gin-example
// @license.name MIT
// @license.url https://Pay/blob/master/LICENSE
// @securityDefinitions.apikey Token
// @in header
// @name token
// @BasePath /
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
