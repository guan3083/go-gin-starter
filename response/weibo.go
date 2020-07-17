package response

type WeiboInfo struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	HotNum string `json:"hot_num"`
	Type   string `json:"type"`
	Url    string `json:"url"`
}
