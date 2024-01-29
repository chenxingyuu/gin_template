package response

const (
	StatusOK = 0 // 请求正常

	// 1xxx 通用错误

	StatusInvalidParams      = 1001 // 参数错误
	StatusUnauthorized       = 1002 // 未授权
	StatusNotFound           = 1003 // 未找到
	StatusInternalError      = 1004 // 内部错误
	StatusForbidden          = 1005 // 禁止访问
	StatusInvalidRequest     = 1006 // 无效的请求
	StatusDuplicateEntry     = 1007 // 重复的记录
	StatusTimeout            = 1008 // 请求超时
	StatusServiceUnavailable = 1009 // 服务不可用

	// 2xxx 用户错误

	StatusUserNotFound   = 2001 // 用户不存在
	StatusUserExists     = 2002 // 用户已存在
	StatusUserWrongPwd   = 2003 // 用户密码错误
	StatusEmailExists    = 2004 // 邮箱已存在
	StatusUsernameExists = 2005 // 用户名已存在

	// 3xxx 业务错误

	StatusMissingToken     = 3001 // 缺少token
	StatusInvalidToken     = 3002 // 无效的token
	StatusExpiredToken     = 3003 // 过期的token
	StatusInvalidCode      = 3004 // 无效的验证码
	StatusInvalidFileType  = 3005 // 无效的文件类型
	StatusFileTooLarge     = 3006 // 文件过大
	StatusOperationFailed  = 3007 // 操作失败
	StatusInvalidInput     = 3008 // 无效的输入
	StatusResourceConflict = 3009 // 资源冲突

	// 4xxx 第三方错误

	StatusThirdPartyError  = 4001 // 第三方错误
	StatusInvalidAPIKey    = 4002 // 无效的 API Key
	StatusAPILimitExceeded = 4003 // API 请求限制超出

	// 5xxx 安全错误

	StatusCSRFMismatch      = 5001 // CSRF Token 不匹配
	StatusSecurityViolation = 5002 // 安全违规

)

func Message(code int) string {
	switch code {
	// 请求正常
	case StatusOK:
		return "ok"

	// 1xxx 通用错误
	case StatusInvalidParams:
		return "invalid params"
	case StatusUnauthorized:
		return "unauthorized"
	case StatusNotFound:
		return "not found"
	case StatusInternalError:
		return "internal error"
	case StatusForbidden:
		return "forbidden"

	// 2xxx 用户错误
	case StatusUserNotFound:
		return "user not found"
	case StatusUserExists:
		return "user exists"
	case StatusUserWrongPwd:
		return "wrong password"

	// 3xxx 业务错误
	case StatusMissingToken:
		return "missing token"
	case StatusInvalidToken:
		return "invalid token"
	case StatusExpiredToken:
		return "expired token"
	case StatusInvalidCode:
		return "invalid code"

	// 4xxx 第三方错误
	case StatusThirdPartyError:
		return "third party error"
	}

	return "unknown error"
}
