package form
//validator验证 但是有自定义翻译器的问题
type SendSmsForm struct{
	Mobile string `form:"mobile" json:"mobile" binging:"required,mobile"`//手机号码有迹可循 自定义validator
	//Type string   `form:"type" json:"type" binging:"required,type,oneog=1 2"`
	//1.注册发送验证码和动态验证码登录发送验证码
}
type VerifySmsCodeForm struct {
	Mobile string `json:"mobile" binding:"required,e164"`
	Code   string `json:"code" binding:"required,len=6"`
}
