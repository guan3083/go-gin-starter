package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-gin-starter/pkg/app"
	"go-gin-starter/pkg/e"
	"go-gin-starter/pkg/util"
)

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	ID          int64  `json:"userId"`
	AccountName string `json:"accountName"`
	RoleId      int    `json:"role_id"`
	RoleName    string `json:"role_name"`
	jwt.StandardClaims
}

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, code := checkToken(c)
		if code != e.SUCCESS {
			app.UnauthorizedResp(c, code, "")
			c.Abort()
			return
		}

		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set(util.TokenKey, claims)
	}
}

func checkToken(c *gin.Context) (*util.Claims, int) {
	token := c.GetHeader(util.HeaderToken)
	if token == "" {
		return nil, e.ERROR_AUTH
	}

	claims, err := util.ParseToken(token)
	if err != nil {
		switch err.(*jwt.ValidationError).Errors {
		case jwt.ValidationErrorExpired:
			return nil, e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
		default:
			return nil, e.ERROR_AUTH_CHECK_TOKEN_FAIL
		}
	}

	return claims, e.SUCCESS
}
