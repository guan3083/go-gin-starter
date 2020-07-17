package system

import (
	"github.com/gin-gonic/gin"
	"go-gin-starter/pkg/app"
	"go-gin-starter/pkg/e"
	"go-gin-starter/service/sysytem_service"
	"io/ioutil"
)

// @Summary 初始化数据库
// @Description 初始化数据库
// @Tags 系统设置
// @accept json
// @Produce  json
// @Success 200 {object}  app.Response
// @Failure 500 {object}  app.Response
// @Router /api/v1/system/initScript  [get]
func InitSqlScript(c *gin.Context) {

	file, err := ioutil.ReadFile("script/go-gin-starter.sql")
	if err != nil {
		app.ErrorResp(c, e.ERROR, err.Error())
		return
	}
	err = sysytem_service.Execute(string(file))
	if err != nil {
		app.ErrorResp(c, e.ERROR, err.Error())
		return
	}
	app.SuccessResp(c, string(file))

}
