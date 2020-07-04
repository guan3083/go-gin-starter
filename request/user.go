package request

//登录
type ReqLoginForm struct {
	// 页面大小
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
}

//获取单个用户信息
type ReqGetUserListForm struct {
	// 页面大小
	PageNo   int `form:"page" json:"page"`
	PageSize int `form:"size" json:"size"`
}

//单个Id
type ReqUserIdForm struct {
	// 页面大小
	Id int64 `form:"id" json:"id" binding:"required"`
}

//添加用户
type ReqAddUserForm struct {
	// 用户名
	UserName string `json:"user_name" gorm:"column:user_name" binding:"required"`
	// 登录密码
	Password string `json:"password" gorm:"column:password" binding:"required"`
	// 姓名
	NickName string `json:"nick_name" gorm:"column:nick_name"`
	// 电话号码
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number"`
	// 地址
	Address string `json:"address" gorm:"column:address"`
}

//修改用户
type ReqUpdateUserForm struct {
	// 主键
	Id int64 `json:"id" gorm:"column:id" binding:"required"`
	// 用户名
	UserName string `json:"user_name" gorm:"column:user_name"`
	// 登录密码
	Password string `json:"password" gorm:"column:password"`
	// 姓名
	NickName string `json:"nick_name" gorm:"column:nick_name"`
	// 电话号码
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number"`
	// 地址
	Address string `json:"address" gorm:"column:address"`
	// 账户状态（1：暂停，2：启用）
	Status int `json:"status" gorm:"column:status"`
}
