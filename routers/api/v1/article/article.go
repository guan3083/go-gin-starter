package article

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-starter/models"
	"go-gin-starter/pkg/app"
	"go-gin-starter/pkg/e"
	"go-gin-starter/pkg/setting"
	"go-gin-starter/pkg/util"
	"go-gin-starter/request"
	"go-gin-starter/response"
	"go-gin-starter/service/article_service"
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
	content, err := crawler_service.Execute(form.Type, form.Uri)
	if err != nil {
		app.ErrorResp(c, e.INVALID_PARAMS, err.Error())
		return
	}
	app.SuccessResp(c, content)
}

// @Summary 文章列表
// @Description 文章查询
// @Tags 文章
// @accept json
// @Produce  json
// @Param page query int false "当前页面"
// @Param size query int false "页数量"
// @Success 200 {object}  app.Response
// @Failure 500 {object}  app.Response
// @Router /api/v1/article/list  [get]
func GetAll(c *gin.Context) {
	var (
		form request.ReqCommonPage
	)
	err := app.BindAndValid(c, &form)
	if err != nil {
		app.ErrorResp(c, e.INVALID_PARAMS, err.Error())
		return
	}

	// 分页查询所有
	pageInfo := util.GetPaginationByCommon(form)
	articleService := article_service.Article{
		Page: pageInfo,
	}

	total, err := articleService.Count()
	if err != nil {
		app.ErrorResp(c, e.ERROR, err.Error())
		return
	}

	data, err := articleService.GetAll()
	if err != nil {
		app.ErrorResp(c, e.ERROR, err.Error())
		return
	}

	app.SuccessResp(c, response.RespArticleInfoList{
		Total: total,
		List:  data,
	})
}

// @Summary 文章列表
// @Description 文章查询
// @Tags 文章
// @accept json
// @Produce  json
// @Param page query int false "当前页面"
// @Param size query int false "页数量"
// @Success 200 {object}  app.Response
// @Failure 500 {object}  app.Response
// @Router /api/v1/article/weibo  [get]
func GetWeiboAll(c *gin.Context) {
	var (
		form request.ReqCommonPage
	)
	err := app.BindAndValid(c, &form)
	if err != nil {
		app.ErrorResp(c, e.INVALID_PARAMS, err.Error())
		return
	}

	content, err := crawler_service.Execute("weibo", setting.AppSetting.WeiboIndex)
	weiboInfos := content.([]response.WeiboInfo)
	total := int64(len(weiboInfos))
	var articles []*models.Article
	//pageNo := form.PageNo - 1 <= 0 ? 1 : form.PageNo
	var pageNo = 1
	if form.PageNo-1 > 0 {
		pageNo = form.PageNo
	}
	start := (pageNo - 1) * form.PageSize
	end := pageNo * form.PageSize
	for i := start; i < end; i++ {
		if int64(i) >= total {
			break
		}
		item := weiboInfos[i]
		contentHtml := fmt.Sprintf("<iframe src='%v' />", item.Url)
		article := &models.Article{
			Id:          int64(i + 1),
			Title:       item.Title,
			ContentHtml: contentHtml,
			HotNum:      item.HotNum,
		}
		articles = append(articles, article)
	}
	app.SuccessResp(c, response.RespArticleInfoList{
		Total: total,
		List:  articles,
	})
}
