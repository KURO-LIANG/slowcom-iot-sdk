package serror

type Error struct {
	Code    int
	Message string
}

func (e Error) Error() string {
	return e.Message
}

var (
	ErrIs系统异常 = &Error{
		Code:    101,
		Message: "System Exception",
	}

	ErrIs请求状态异常 = &Error{
		Code:    102,
		Message: "Http Exception",
	}

	ErrIs数据解析异常 = &Error{
		Code:    103,
		Message: "Data Parsing Exception",
	}
)

// New 业务异常
func New(code int, message string) error {
	return &Error{
		Code:    code,
		Message: message,
	}
}
