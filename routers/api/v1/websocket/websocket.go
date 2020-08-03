package websocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-starter/pkg/app"
	"go-gin-starter/pkg/e"
	"go-gin-starter/pkg/util"
	"go-gin-starter/request"
)

//@Summary 发送消息
//@Description 发送消息
//@Tags websocket Test
//@accept json
//@Produce  json
//@Param page query int false "当前页面"
//@Param size query int false "页数量"
//@Success 200 {object}  app.Response
//@Failure 500 {object}  app.Response
//@Router /api/v1/ws/send  [get]
func Send(c *gin.Context) {
	var (
		form request.ReqCommonPage
	)
	err := app.BindAndValid(c, &form)
	if err != nil {
		app.ErrorResp(c, e.INVALID_PARAMS, err.Error())
		return
	}
	res := fmt.Sprintf("-%v-%v-", form.PageNo, form.PageSize)
	util.Send("admin", res)
	app.SuccessResp(c, res)
}
