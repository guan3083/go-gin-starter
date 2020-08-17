package shoes

import (
	"github.com/gin-gonic/gin"
	"go-gin-starter/pkg/app"
	"go-gin-starter/pkg/e"
	"go-gin-starter/pkg/util"
	"go-gin-starter/request"
	"go-gin-starter/response"
	"go-gin-starter/service/shoes_service"
)

// @Summary shoes信息
// @Description
// @Tags shoes
// @accept json
// @Produce  json
// @Param page query int false "当前页面"
// @Param size query int false "页数量"
// @Success 200 {object}  app.Response
// @Failure 500 {object}  app.Response
// @Router /api/v1/shoes/list  [get]
func GetShoesList(c *gin.Context) {
	var (
		form request.ReqGetUserListForm
	)
	err := app.BindAndValid(c, &form)
	if err != nil {
		app.ErrorResp(c, e.INVALID_PARAMS, err.Error())
		return
	}

	offset, limit := util.GetPaginationParams(form.PageNo, form.PageSize)

	shoesService := shoes_service.Shoes{
		PageNum:  offset,
		PageSize: limit,
	}
	total, err := shoesService.GetTotals()
	if err != nil {
		app.ErrorResp(c, e.ERROR, err.Error())
		return
	}

	shoes, err := shoesService.GetAll()
	if err != nil {
		app.ErrorResp(c, e.ERROR, err.Error())
		return
	}
	app.SuccessResp(c, response.RespShoesList{
		Total: total,
		List:  shoes,
	})
}
