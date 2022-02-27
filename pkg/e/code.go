package e

const (
	// 通用错误
	Success       = 200
	Error         = 500
	InvalidParams = 400

	// 成员错误
	ErrorNotExistUser  = 10001
	ErrorHaveExistUser = 10002

	// token 错误
	ErrorUserCheckTokenFail    = 30001 //token 错误
	ErrorUserCheckTokenTimeout = 30002 //token 过期
	ErrorUserToken             = 30003 //token 生成失败

	// 数据库错误
	ErrorDatabase = 40001
)
