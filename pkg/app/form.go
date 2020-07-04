package app

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-gin-starter/pkg/errors"
	"go-gin-starter/pkg/logf"
	vali "go-gin-starter/pkg/validation"
	"net/http"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) error {
	var err error
	if c.Request.Method == http.MethodGet {
		err = c.ShouldBindQuery(form)
	} else if c.Request.Method == http.MethodPost {
		err = c.ShouldBindJSON(form)
	}

	var errStr string
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errStr = vali.TranslateOneError(err.(validator.ValidationErrors))
		case *json.UnmarshalTypeError:
			unmarshalTypeError := err.(*json.UnmarshalTypeError)
			errStr = fmt.Errorf("%s 类型错误，期望类型 %s", unmarshalTypeError.Field, unmarshalTypeError.Type.String()).Error()
		default:
			errStr = err.Error()
		}
		logf.Error(errStr)
		return errors.New(errStr)
	}

	return nil
}
