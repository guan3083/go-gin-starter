package user

import (
	"github.com/gin-gonic/gin"
	"go-gin-starter/pkg/app"
	"go-gin-starter/pkg/e"
	"go-gin-starter/pkg/setting"
	"go-gin-starter/pkg/util"
	"go-gin-starter/request"
	"go-gin-starter/response"
	"go-gin-starter/service/user_service"
)

// @Summary 登录
// @Description 登录
// @Tags 用户
// @accept json
// @Produce  json
// @Param user body request.ReqLoginForm true "登录"
// @Success 200 {object}  app.Response
// @Failure 500 {object}  app.Response
// @Router /api/v1/login  [post]
func Login(c *gin.Context) {
	var (
		form request.ReqLoginForm
	)
	err := app.BindAndValid(c, &form)
	if err != nil {
		app.ErrorResp(c, e.INVALID_PARAMS, err.Error())
		return
	}
	userServer := user_service.User{}
	user, err := userServer.GetByUserName(form.UserName)
	if err != nil {
		app.ErrorResp(c, e.ERROR, "账户或者密码不存在")
		return
	}
	if user == nil {
		app.ErrorResp(c, e.ERROR, "用户不存在")
		return
	}
	salt := util.GetMd5String(form.UserName)
	md5Equals := util.MD5Equals(form.Password, salt, user.Password)
	if !md5Equals {
		app.ErrorResp(c, e.ERROR, "密码错误")
		return
	}
	// 2、token生成
	token, err := util.GenerateToken(form.UserName, user.NickName)
	if err != nil {
		app.ErrorResp(c, e.ERROR, err.Error())
		return
	}
	app.SuccessResp(c, token)

}

// @Summary 注册
// @Description 添加用户信息
// @Tags 用户
// @accept json
// @Produce  json
// @Param form body request.ReqAddUserForm true "reqBody"
// @Success 200 {object}  app.Response
// @Failure 500 {object}  app.Response
// @Router /api/v1/register  [post]
func RegisterUser(c *gin.Context) {
	var (
		form request.ReqAddUserForm
	)
	err := app.BindAndValid(c, &form)
	if err != nil {
		app.ErrorResp(c, e.INVALID_PARAMS, err.Error())
		return
	}
	salt := util.GetMd5String(form.UserName)
	password := util.EncodeMD5(form.Password, salt)

	userServer := user_service.User{
		UserName:    form.UserName,
		Password:    password,
		SecretKey:   salt,
		NickName:    form.NickName,
		PhoneNumber: form.PhoneNumber,
		Address:     form.Address,
	}
	err = userServer.AddUser()
	if err != nil {
		app.ErrorResp(c, e.ERROR, err.Error())
		return
	}
	app.SuccessResp(c, nil)
}

// @Summary 获取用户信息
// @Description 获取用户信息
// @Tags 用户
// @accept json
// @Produce  json
// @Param page query int false "当前页面"
// @Param size query int false "页数量"
// @Success 200 {object}  app.Response
// @Failure 500 {object}  app.Response
// @Router /api/v1/user/list  [get]
// @Security Token
func GetUserList(c *gin.Context) {
	var (
		form request.ReqGetUserListForm
	)
	err := app.BindAndValid(c, &form)
	if err != nil {
		app.ErrorResp(c, e.INVALID_PARAMS, err.Error())
		return
	}

	offset, limit := util.GetPaginationParams(form.PageNo, form.PageSize)

	userServer := user_service.User{
		PageNum:  offset,
		PageSize: limit,
	}
	total, err := userServer.Count()
	if err != nil {
		app.ErrorResp(c, e.ERROR, err.Error())
		return
	}

	users, err := userServer.GetAll()
	if err != nil {
		app.ErrorResp(c, e.ERROR, err.Error())
		return
	}
	app.SuccessResp(c, response.RespUserInfoList{
		Total: total,
		List:  users,
	})
}

// @Summary 获取单个用户信息
// @Description 获取用户信息
// @Tags 用户
// @accept json
// @Produce  json
// @Param id query int64 true "用户ID"
// @Success 200 {object}  app.Response
// @Failure 500 {object}  app.Response
// @Router /api/v1/user/get  [get]
// @Security Token
func GetUser(c *gin.Context) {
	var (
		form request.ReqUserIdForm
	)
	err := app.BindAndValid(c, &form)
	if err != nil {
		app.ErrorResp(c, e.INVALID_PARAMS, err.Error())
		return
	}
	userServer := user_service.User{}
	user, err := userServer.GetById(form.Id)
	if err != nil {
		app.ErrorResp(c, e.ERROR, err.Error())
		return
	}
	app.SuccessResp(c, user)
}

// @Summary 修改用户
// @Description 修改用户信息
// @Tags 用户
// @accept json
// @Produce  json
// @Param form body request.ReqUpdateUserForm true "reqBody"
// @Success 200 {object}  app.Response
// @Failure 500 {object}  app.Response
// @Router /api/v1/user/update  [post]
// @Security Token
func UpdateUser(c *gin.Context) {
	var (
		form request.ReqUpdateUserForm
	)
	err := app.BindAndValid(c, &form)
	if err != nil {
		app.ErrorResp(c, e.INVALID_PARAMS, err.Error())
		return
	}
	userServer := user_service.User{
		Id:          form.Id,
		UserName:    form.UserName,
		Password:    form.Password,
		SecretKey:   setting.AppSetting.MD5Salt,
		NickName:    form.NickName,
		PhoneNumber: form.PhoneNumber,
		Address:     form.Address,
		Status:      form.Status,
	}
	err = userServer.UpdateUser()
	if err != nil {
		app.ErrorResp(c, e.ERROR, err.Error())
		return
	}
	app.SuccessResp(c, nil)
}

// @Summary 删除用户
// @Description 删除用户信息
// @Tags 用户
// @accept json
// @Produce  json
// @Param form body request.ReqUserIdForm  true "Id"
// @Success 200 {object}  app.Response
// @Failure 500 {object}  app.Response
// @Router /api/v1/user/delete  [post]
// @Security Token
func DeleteUser(c *gin.Context) {
	var (
		form request.ReqUserIdForm
	)
	err := app.BindAndValid(c, &form)
	if err != nil {
		app.ErrorResp(c, e.INVALID_PARAMS, err.Error())
		return
	}
	userServer := user_service.User{}
	err = userServer.DeleteUser(form.Id)
	if err != nil {
		app.ErrorResp(c, e.ERROR, err.Error())
		return
	}
	app.SuccessResp(c, nil)
}
