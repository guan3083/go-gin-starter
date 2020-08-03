package crawler_service

import (
	"fmt"
	htmlquery "github.com/antchfx/xquery/html"
	"go-gin-starter/pkg/setting"
	"go-gin-starter/response"
	"net/http"
	"strings"
)

func Execute(source, url string) (interface{}, error) {

	switch source {
	case "csdn":
		return executeCsdn(url)
	case "weibo":
		return executeWeibo(url)
	}
	return "", nil
}

func executeCsdn(url string) (string, error) {
	get, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer get.Body.Close()
	root, _ := htmlquery.Parse(get.Body)
	titleArticle := htmlquery.Find(root, "//*[@class='title-article']")
	if len(titleArticle) == 0 {
		return fmt.Sprintf("地址%v为空", url), nil
	}
	var title string
	for _, row := range titleArticle {
		item := htmlquery.Find(row, ".")
		title = htmlquery.InnerText(item[0])
	}
	contentArticle := htmlquery.Find(root, "//*[@id='content_views']")
	content := htmlquery.OutputHTML(contentArticle[0], true)
	return strings.ReplaceAll(fmt.Sprintf("<h1>%v</h1> %v", title, content), "\n", ""), nil
}

func executeWeibo(url string) ([]response.WeiboInfo, error) {
	get, err := http.Get(url)
	if err != nil {
		return nil, nil
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
	return proxies, nil
}
