package e

const (
	// 通用错误
	Success       = 200
	Error         = 500
	InvalidParams = 400

	// 成员错误
	ErrorNotExistUser        = 10001
	ErrorHaveExistUser       = 10002
	ErrorFailEncryption      = 10003
	ErrorPasswordFailCompare = 10004
	ErrorNewPasswordNull     = 10005

	// token 错误
	ErrorUserCheckTokenFail    = 30001 //token 错误
	ErrorUserCheckTokenTimeout = 30002 //token 过期
	ErrorUserToken             = 30003 //token 生成失败

	// JSON 错误
	ErrorJSONNotMatch = 40001 // JSON类型不匹配

	// 数据库错误
	ErrorDatabase = 50001
)
