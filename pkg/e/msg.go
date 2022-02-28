package e

var MsgFlags = map[int]string{
	Success:       "ok",
	Error:         "fail",
	InvalidParams: "请求参数错误",

	ErrorNotExistUser:        "用户不存在",
	ErrorHaveExistUser:       "用户已存在",
	ErrorFailEncryption:      "加密失败",
	ErrorPasswordFailCompare: "用户密码错误",

	ErrorUserCheckTokenFail:    "Token 鉴权失败",
	ErrorUserCheckTokenTimeout: "Token 已超时",
	ErrorUserToken:             "Token 生成失败",

	ErrorDatabase: "数据库操作错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[Error]
}
