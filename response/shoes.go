package response

import "go-gin-starter/models"

//获取用户信息列表
type RespShoesList struct {
	// 总数
	Total int64 `json:"total"`
	// 数据列表
	List []*models.Shoes `json:"list"`
}
