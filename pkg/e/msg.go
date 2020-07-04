package e

var MsgFlags = map[int]string{
	SUCCESS:                        "success",
	ERROR:                          "内部服务异常",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_AUTH:                     "认证失败",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_IP_RISK:                  "请求次数过多，请等待15分钟后再试",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
