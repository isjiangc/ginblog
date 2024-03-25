package handler

import (
	"net/http"
	"time"

	"ginblog/internal/middleware"
	"ginblog/internal/model"
	"ginblog/internal/service"
	"ginblog/pkg/helper/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type LoginHandler interface {
	Login(ctx *gin.Context)
	LoginFront(ctx *gin.Context)
}

type loginHandler struct {
	*Handler
	conf        *viper.Viper
	userService service.UserService
}

func NewLoginHandler(handler *Handler, conf *viper.Viper, userService service.UserService) LoginHandler {
	return &loginHandler{
		Handler:     handler,
		conf:        conf,
		userService: userService,
	}
}

// Login 后台登录
func (l *loginHandler) Login(ctx *gin.Context) {
	var formData model.User
	_ = ctx.ShouldBindJSON(&formData)
	var token string
	var code int
	formData, code = l.userService.CheckLogin(formData.Username, formData.Password)
	if code == errmsg.SUCCESS {
		setToken(ctx, formData, l.conf)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    formData.Username,
			"id":      formData.ID,
			"message": errmsg.GetErrorMsg(code),
			"token":   token,
		})
	}
}

// LoginFront 前台登录
func (l *loginHandler) LoginFront(ctx *gin.Context) {
	var formData model.User
	_ = ctx.ShouldBindJSON(&formData)
	var code int
	formData, code = l.userService.CheckLoginFront(formData.Username, formData.Password)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    formData.Username,
		"id":      formData.ID,
		"message": errmsg.GetErrorMsg(code),
	})
}

func setToken(ctx *gin.Context, user model.User, conf *viper.Viper) {
	j := middleware.NewJWT(conf)
	claims := middleware.MyClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 604800,
			Issuer:    "GinBlog",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR,
			"message": errmsg.GetErrorMsg(errmsg.ERROR),
			"token":   token,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  200,
			"data":    user.Username,
			"id":      user.ID,
			"message": errmsg.GetErrorMsg(200),
			"token":   token,
		})
	}
	return
}
