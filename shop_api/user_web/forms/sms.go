package forms

type SendSmsForm struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"` //和注册验证器的mobile对应
	Type   uint   `form:"type" json:"type" binding:"required,oneof=1 2"`
}
