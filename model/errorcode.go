package model

// 错误码规则:
// (1) 错误码需为 > 0 的数;
//
// (2) 错误码为 5 位数:
//              ----------------------------------------------------------
//                  第1位               2、3位                  4、5位
//              ----------------------------------------------------------
//                服务级错误码          模块级错误码	         具体错误码
//              ----------------------------------------------------------

type ErrorCode struct {
	Err  error
	Code int
	Msg  string
}

// AddErr 添加具体的代码错误
func (e ErrorCode) AddErr(err error) ErrorCode {
	return ErrorCode{
		Code: e.Code,
		Msg:  e.Msg,
		Err:  err,
	}
}

func responseErrCode(code int, msg string) ErrorCode {
	return ErrorCode{
		Code: code,
		Msg:  msg,
	}
}

//ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
//ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
//ERROR_AUTH_TOKEN:               "Token生成失败",
//ERROR_AUTH:                     "Token错误",

var (
	Err                   = responseErrCode(400, "接口错误") // 通用错误
	ErrParam              = responseErrCode(10001, "参数有误")
	ErrLonginParam        = responseErrCode(10001, "用户名或密码错误")
	ErrSignParam          = responseErrCode(10002, "签名参数有误")
	ErrAuthToken          = responseErrCode(10003, "token错误")
	ErrGenToken           = responseErrCode(10004, "token生成错误")
	ErrRegisterParam      = responseErrCode(10005, "用户名已存在")
	ErrAuthCheckTokenFail = responseErrCode(10005, "token鉴权失败")
	ErrDb                 = responseErrCode(20003, "数据库错误")

	// ......
)
