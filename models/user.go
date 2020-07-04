package models

import (
	"github.com/jinzhu/gorm"
	"go-gin-starter/pkg/util"
)

type User struct {

	// 主键id
	Id int64 `json:"id" gorm:"primary_key" gorm:"column:id"`
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
	// token盐值
	SecretKey string `json:"secret_key" gorm:"column:secret_key"`
	// 账户状态（1：暂停，2：启用）
	Status int `json:"status" gorm:"column:status"`
	// 创建时间
	CreateTime util.JSONTime `json:"create_time" gorm:"column:create_time"`
	// 更新时间
	UpdateTime util.JSONTime `json:"update_time" gorm:"column:update_time"`

	Session *Session `json:"-" gorm:"-"`
}

// 设置User的表名为`user`
func (User) TableName() string {
	return "user"
}

func NewUserModel(session *Session) *User {
	return &User{Session: session}
}

// GetUserTotal gets the total number of users based on the constraints
func (a *User) GetUserTotal(maps interface{}) (int64, error) {
	var count int64
	err := a.Session.db.Model(&User{}).Where(maps).Count(&count).Error
	return count, err
}

// GetUsers gets a list of users based on paging constraints
func (a *User) GetUsers(pageNum int, pageSize int, maps interface{}) ([]*User, error) {
	var users []*User
	err := a.Session.db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return users, nil
}

// GetUser Get a single user based on ID
func (a *User) GetUser(id int64) (*User, error) {
	user := new(User)
	err := a.Session.db.Where("id = ?", id).First(user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	} else if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return user, nil
}

func (a *User) GetUserByName(username string) (*User, error) {
	user := new(User)
	err := a.Session.db.Where("user_name = ? ", username).First(user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	} else if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return user, nil
}

// AddUser add a single user
func (a *User) AddUser(user *User) error {
	tx := GetSessionTx(a.Session)
	return tx.Create(user).Error
}

func (a *User) UpdateUser(user *User) error {
	tx := GetSessionTx(a.Session)
	return tx.Model(&User{}).Updates(user).Error
}

// DeleteUser delete a single user
func (a *User) DeleteUser(id int64) error {
	tx := GetSessionTx(a.Session)
	return tx.Where("id = ? ", id).Delete(User{}).Error
}
