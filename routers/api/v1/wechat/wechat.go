package wechat

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-starter/pkg/app"
	"go-gin-starter/pkg/e"
)

//
type ReqGetSessionForm struct {
	//
	Code string `form:"code" json:"code" binding:"required"`
}

func GetSessionKey(c *gin.Context) {
	var (
		form ReqGetSessionForm
	)
	err := app.BindAndValid(c, &form)
	if err != nil {
		app.ErrorResp(c, e.INVALID_PARAMS, err.Error())
		return
	}
	fmt.Println("请求的code:", form.Code)
}
