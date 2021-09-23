package util

const (
	ErrCodeSuccess           = 0    // 成功
	ErrCodeParameter         = 1001 // 参数错误
	ErrCodeUserExist         = 1002 // 用户已经存在
	ErrCodeServerBusy        = 1003 // 服务器繁忙
	ErrCodeUserNotExist      = 1004 // 用户不存在
	ErrCodeUserPasswordWrong = 1005 // 账号或密码错误
	ErrCodeCaptionHit        = 1006 // 标题中含有非法内容
	ErrCodeContentHit        = 1007 // 内容中含有非法内容
)

func GetMessage(code int) (message string) {
	switch code {
	case ErrCodeSuccess:
		message = "success"
	case ErrCodeParameter:
		message = "参数错误"
	case ErrCodeUserExist:
		message = "用户已经存在"
	case ErrCodeServerBusy:
		message = "服务器繁忙"
	case ErrCodeUserNotExist:
		message = "用户不存在"
	case ErrCodeUserPasswordWrong:
		message = "账号或密码错误"
	case ErrCodeCaptionHit:
		message = "标题中含有非法内容"
	case ErrCodeContentHit:
		message = "内容中含有非法内容"
	default:
		message = "未知错误"
	}
	return
}
