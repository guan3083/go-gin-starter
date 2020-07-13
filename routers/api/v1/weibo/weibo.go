package weibo

import (
	"fmt"
	htmlquery "github.com/antchfx/xquery/html"
	"github.com/gin-gonic/gin"
	"go-gin-starter/pkg/app"
	"go-gin-starter/pkg/e"
	"go-gin-starter/pkg/setting"
	"go-gin-starter/response"
	"net/http"
	"strings"
)

// @Summary 获取微博热点数据
// @Description 获取微博热点数据
// @Tags 通用
// @accept json
// @Produce  json
// @Success 200 {object}  app.Response
// @Failure 500 {object}  app.Response
// @Router /api/v1/general/weibo/list  [get]
func GetWeiboHotData(c *gin.Context) {
	get, err := http.Get(setting.AppSetting.WeiboIndex)
	if err != nil {
		fmt.Println("err:", err)
		app.ErrorResp(c, e.ERROR, err.Error())
		return
	}
	defer get.Body.Close()

	root, _ := htmlquery.Parse(get.Body)
	var proxies []response.WeiboInfo
	tr := htmlquery.Find(root, "//*[@id='pl_top_realtimehot']/table/tbody/tr")
	for i, row := range tr {
		item := htmlquery.Find(row, "./td")
		port := htmlquery.InnerText(item[1])
		type_ := htmlquery.InnerText(item[2])
		split := strings.Split(port, "\n")
		url := setting.AppSetting.WeiboSearch
		info := response.WeiboInfo{
			Id:     i,
			Title:  strings.TrimSpace(split[1]),
			HotNum: strings.TrimSpace(split[2]),
			Type:   strings.TrimSpace(type_),
			Url:    strings.ReplaceAll(url, "${}", strings.TrimSpace(split[1])),
		}
		proxies = append(proxies, info)
	}
	app.SuccessResp(c, proxies)

}
