package dto

import (
	"github.com/e421083458/go_gateway/public"
	"github.com/gin-gonic/gin"
)

// 管理员登陆请求结构体
type AdminLoginInput struct {
	UserName string `json:"username" form:"username" comment:"管理员用户名" example:"admin" validate:"required,valid_username"` //管理员用户名
	Password string `json:"password" form:"password" comment:"密码" example:"123456" validate:"required"`                   //密码
}

func (param *AdminLoginInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

// 管理员登陆回复结构体
type AdminLoginOutput struct {
	Token string `json:"token" form:"token" comment:"token" example:"token" validate:""` //token
}
