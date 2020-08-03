package article_service

import (
	"go-gin-starter/models"
	"go-gin-starter/pkg/util"
	"go-gin-starter/request"
)

type Article struct {
	Id int64
	// 标题
	Title string
	// 内容
	Content string
	// html格式的内容
	ContentHtml string
	// 上架时间
	UpDate string
	// 来源
	Source string
	// 作者
	Author string
	// 创建时间
	CreateTime util.JSONTime
	// 更新时间
	UpdateTime util.JSONTime

	// 分页
	Page request.ReqCommonPage
}

func (a *Article) GetAll() ([]*models.Article, error) {
	session := models.NewSession()
	all, err := models.NewArticleModel(session).GetAll(a.Page.PageNo, a.Page.PageSize)
	return all, err
}

func (a *Article) Count() (int64, error) {
	session := models.NewSession()
	return models.NewArticleModel(session).GetArticleTotal()
}
