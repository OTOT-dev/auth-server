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

var (
	Err            = responseErrCode(400, "接口错误") // 通用错误
	ErrParam       = responseErrCode(10001, "参数有误")
	ErrSignParam   = responseErrCode(10002, "签名参数有误")
	ErrDb          = responseErrCode(10003, "数据库错误")
	ErrUserService = responseErrCode(20100, "用户服务异常")
	ErrUserPhone   = responseErrCode(20101, "用户手机号不合法")
	ErrUserCaptcha = responseErrCode(20102, "用户验证码有误")

	// ......
)
