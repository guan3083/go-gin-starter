package article

import (
	"github.com/gin-gonic/gin"
	"go-gin-starter/pkg/app"
	"go-gin-starter/pkg/e"
	"go-gin-starter/request"
	"go-gin-starter/service/crawler_service"
)

// @Summary 文章爬取
// @Description 文章爬取
// @Tags 文章
// @accept json
// @Produce  json
// @Param form body request.ReqCrawler true "reqBody"
// @Success 200 {object}  app.Response
// @Failure 500 {object}  app.Response
// @Router /api/v1/article/crawler  [post]
func Crawler(c *gin.Context) {
	var (
		form request.ReqCrawler
	)
	err := app.BindAndValid(c, &form)
	if err != nil {
		app.ErrorResp(c, e.INVALID_PARAMS, err.Error())
		return
	}
	content, err := crawler_service.Execute(form.Uri)
	if err != nil {
		app.ErrorResp(c, e.INVALID_PARAMS, err.Error())
		return
	}
	app.SuccessResp(c, content)
}
