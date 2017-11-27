package define

const (
	UnAuthorizedError = "未授权"
	ParameterError = "参数错误"
	UserExistError = "用户已存在"
	UserNotExistError = "用户不存在"
	InvalidEmailError = "无效的邮箱"
	WrongPasswordError = "账号或密码错误"
	InactiveError = "账号未激活"
	InvalidCodeError = "验证码错误"
)


const (

	RegisterSuccess = "用户注册成功"
	RegistEmailSuccess = "注册成功,请前往邮箱激活账号"
	LoginSuccess = "登录成功"
	ActivateSuccess = "激活成功"
	MessageSendSuccess = "验证码发送成功"
	UploadSuccess = "上传成功"
)

const (
	EmailPattern = `^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`
	PhonePattern = `^1\d{10}$`
)



