package crawler_service

import (
	"fmt"
	htmlquery "github.com/antchfx/xquery/html"
	"net/http"
	"strings"
)

func Execute(url string) (string, error) {

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
