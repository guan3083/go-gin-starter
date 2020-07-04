package user_service

import (
	"go-gin-starter/models"
	"go-gin-starter/pkg/util"
)

const (
	userStatusForbidden = 1
	userStatusNormal    = 2
)

type User struct {
	// 主键id
	Id int64
	// 用户名
	UserName string
	// 登录密码
	Password string
	// 交易密码
	TradeKey string
	// 姓名
	NickName string
	// 电话号码
	PhoneNumber string
	// 地址
	Address string
	// token盐值
	SecretKey string
	// 账户状态（1：暂停，2：启用）
	Status int
	// 创建时间
	CreateTime util.JSONTime
	// 更新时间
	UpdateTime util.JSONTime

	PageNum  int
	PageSize int
}

func (a *User) GetAll() ([]*models.User, error) {
	session := models.NewSession()
	users, err := models.NewUserModel(session).GetUsers(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (a *User) GetById(id int64) (*models.User, error) {
	session := models.NewSession()
	users, err := models.NewUserModel(session).GetUser(id)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (a *User) GetByUserName(username string) (*models.User, error) {
	session := models.NewSession()
	users, err := models.NewUserModel(session).GetUserByName(username)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (a *User) Count() (int64, error) {
	session := models.NewSession()
	return models.NewUserModel(session).GetUserTotal(a.getMaps())
}

func (a *User) AddUser() error {
	session := models.NewSession()
	user := &models.User{
		UserName:    a.UserName,
		Password:    a.Password,
		NickName:    a.NickName,
		SecretKey:   a.SecretKey,
		PhoneNumber: a.PhoneNumber,
		Address:     a.Address,
		Status:      userStatusNormal,
	}
	err := models.NewUserModel(session).AddUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (a *User) UpdateUser() error {
	session := models.NewSession()
	user := &models.User{
		Id:          a.Id,
		UserName:    a.UserName,
		Password:    a.Password,
		NickName:    a.NickName,
		SecretKey:   a.SecretKey,
		PhoneNumber: a.PhoneNumber,
		Address:     a.Address,
		Status:      userStatusNormal,
	}
	err := models.NewUserModel(session).UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (a *User) DeleteUser(id int64) error {
	session := models.NewSession()
	return models.NewUserModel(session).DeleteUser(id)
}

func (a *User) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	if a.UserName != "" {
		maps["user_name"] = a.UserName
	}
	if a.Password != "" {
		maps["password"] = a.Password
	}
	if a.TradeKey != "" {
		maps["trade_key"] = a.TradeKey
	}

	if a.NickName != "" {
		maps["nick_name"] = a.NickName
	}
	if a.PhoneNumber != "" {
		maps["phone_number"] = a.PhoneNumber
	}
	if a.Address != "" {
		maps["address"] = a.Address
	}
	if a.SecretKey != "" {
		maps["secret_key"] = a.SecretKey
	}
	if a.Status != 0 {
		maps["status"] = a.Status
	}

	return maps
}
