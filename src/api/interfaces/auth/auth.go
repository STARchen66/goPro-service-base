package auth

import "github.com/gin-gonic/gin"

type AauthHandlerInterface interface {

	// Login 登录
	Login(cxt *gin.Context)

	// Registered 注册
	Registered(cxt *gin.Context)

	// SendEmailCode 发送邮件
	SendEmailCode(cxt *gin.Context)
}
