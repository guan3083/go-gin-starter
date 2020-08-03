package response

import "go-gin-starter/models"

//
type RespArticleInfoList struct {
	// 总数
	Total int64 `json:"total"`
	// 数据列表
	List []*models.Article `json:"list"`
}
