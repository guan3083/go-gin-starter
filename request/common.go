package request

//获取单个用户信息
type ReqCommonPage struct {
	// 页面大小
	PageNo   int `form:"page" json:"page"`
	PageSize int `form:"size" json:"size"`
}
