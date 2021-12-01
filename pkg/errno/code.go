package errno

// 统一存自定义的错误码
var (
	// 定义通用错误
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "参数跟struct对不上"}

	ErrValidation = &Errno{Code: 20001, Message: "校验错误"}
	ErrDatabase   = &Errno{Code: 20002, Message: "数据库错误"}
	ErrToken      = &Errno{Code: 20003, Message: "token错误"}

	// 定义业务错误
	// 用户不存在错误
	ErrUserNotFound      = &Errno{Code: 20102, Message: "用户不存在"}
	ErrEncrypt           = &Errno{Code: 20101, Message: "用户密码解密错误"}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "token无效"}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "密码错误"}
)
