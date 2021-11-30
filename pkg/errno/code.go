package errno

// 统一存自定义的错误码
var (
	// 定义通用错误
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	// 定义业务错误
	// 用户不存在错误
	ErrUserNotFound = &Errno{Code: 20102, Message: "用户不存在"}
)
