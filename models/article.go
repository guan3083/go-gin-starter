package models

import "go-gin-starter/pkg/util"

type Article struct {
	Id int64 `json:"id" gorm:"column:id"`
	// 标题
	Title string `json:"title" gorm:"column:title"`
	// 内容
	Content string `json:"content" gorm:"column:content"`
	// html格式的内容
	ContentHtml string `json:"content_html" gorm:"column:content_html"`
	// 上架时间
	UpDate string `json:"up_date" gorm:"column:up_date"`
	// 来源
	Source string `json:"source" gorm:"column:source"`
	// 作者
	Author string `json:"author" gorm:"column:author"`
	// 热度
	HotNum string `json:"hot_num" gorm:"column:hot_num"`
	// 创建时间
	CreateTime util.JSONTime `json:"create_time" gorm:"column:create_time"`
	// 更新时间
	UpdateTime util.JSONTime `json:"update_time" gorm:"column:update_time"`

	Session *Session `json:"-" gorm:"-"`
}

// 设置Article的表名为`article`
func (Article) TableName() string {
	return "article"
}

func NewArticleModel(session *Session) *Article {
	return &Article{Session: session}
}
func (a *Article) GetAll(page, size int) ([]*Article, error) {
	var articles []*Article
	err := a.Session.db.Offset(page).Limit(size).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (a *Article) GetArticleTotal() (int64, error) {
	var count int64
	err := a.Session.db.Model(&Article{}).Count(&count).Error
	return count, err
}
