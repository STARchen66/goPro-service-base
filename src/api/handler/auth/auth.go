package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"goImPro-service/pkg/hash"
	"goImPro-service/pkg/jwt"
	"goImPro-service/pkg/mysql"
	"goImPro-service/pkg/response"
	"goImPro-service/src/api/requests"
	"goImPro-service/src/api/services"
	"goImPro-service/src/config"
	"goImPro-service/src/dao/auth_dao"
	"goImPro-service/src/enum"
	"goImPro-service/src/helpers"
	"goImPro-service/src/models/user"
	"net/http"
	"time"
)

type AuthHandler struct {
}

type loginResponse struct {
	ID         int64  `json:"id"`          //用户id
	UID        string `json:"uid"`         // uid
	Name       string `json:"name"`        //名称
	Avatar     string `json:"avatar"`      //头像
	Email      string `json:"email"`       //邮箱账号
	Token      string `json:"token"`       // token
	ExpireTime int64  `json:"expire_time"` // token过期时间
	Ttl        int64  `json:"ttl"`         // token有效期
}

var (
	auth auth_dao.AuthDao
)

func (*AuthHandler) Login(ctx *gin.Context) {
	params := requests.LoginForm{
		Email:    ctx.PostForm("email"),
		Password: ctx.PostForm("password"),
	}

	errs := validator.New().Struct(params)
	if errs != nil {
		response.FailResponse(http.StatusInternalServerError, errs.Error()).WriteTo(ctx)
		return
	}

	var users user.BaseUsers

	res := mysql.DB.Table("base_users").Where("email=?", params.Email).First(&users)

	if res.RowsAffected == 0 {
		response.FailResponse(http.StatusInternalServerError, "邮箱未注册").ToJson(ctx)
		return
	}

	fmt.Println(users.Password)
	if !hash.BcryptCheck(params.Password, users.Password) {
		response.FailResponse(http.StatusInternalServerError, "密码错误").ToJson(ctx)
		return
	}

	ttl := config.Conf.JWT.Ttl
	expireAtTime := time.Now().Unix() + ttl
	token := jwt.NewJWT().IssueToken(
		users.ID,
		users.Uid,
		users.Name,
		users.Email,
		expireAtTime,
	)

	response.SuccessResponse(&loginResponse{
		ID:         users.ID,
		UID:        users.Uid,
		Name:       users.Name,
		Avatar:     users.Avatar,
		Email:      users.Email,
		ExpireTime: expireAtTime,
		Token:      token,
		Ttl:        ttl,
	}).WriteTo(ctx)

	return

}

func (*AuthHandler) Registered(ctx *gin.Context) {

	params := requests.RegisteredForm{
		Email:          ctx.PostForm("email"),
		Name:           ctx.PostForm("name"),
		EmailType:      helpers.StringToInt(ctx.DefaultPostForm("email_type", "1")),
		Password:       ctx.PostForm("password"),
		PasswordRepeat: ctx.PostForm("password_repeat"),
		Code:           ctx.PostForm("code"),
	}

	err := validator.New().Struct(params)

	if err != nil {
		response.FailResponse(enum.ParamError, err.Error()).WriteTo(ctx)
		return
	}

	ok, filed := user.IsUserExits(params.Email, params.Name)

	if ok {
		response.FailResponse(enum.ParamError, fmt.Sprintf("%s已经存在了", filed)).WriteTo(ctx)
		return
	}

	if config.Conf.Server.Mode == "release" {
		var emailService services.EmailService

		if !emailService.CheckCode(params.Email, params.Code, params.EmailType) {
			response.FailResponse(enum.ParamError, "邮件验证码不正确").WriteTo(ctx)
			return
		}
	}

	auth.CreateUser(params.Email, params.Password, params.Name)

	// 投递消息
	//services.InitChatBotMessage(1, id)

	response.SuccessResponse().ToJson(ctx)
	return
}

func (*AuthHandler) SendEmailCode(ctx *gin.Context) {
}
