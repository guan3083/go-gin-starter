package request

type ReqCrawler struct {
	Type string `json:"type" binding:"required,oneof=csdn"`
	Uri  string `json:"uri"  binding:"required,url"`
}
