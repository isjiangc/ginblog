package middleware

import (
	"net/http"
	"strings"

	"ginblog/pkg/helper/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type JWT struct {
	JwtKey []byte
}

func NewJWT(conf *viper.Viper) *JWT {
	return initJwt(conf)
}

func initJwt(conf *viper.Viper) *JWT {
	return &JWT{[]byte(conf.GetString("security.api_sign.app_key"))}
}

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var (
	TokenExpired     error = errors.New("Token已过期,请重新登录")
	TokenNotValidYet error = errors.New("Token无效,请重新登录")
	TokenMalformed   error = errors.New("Token不正确,请重新登录")
	TokenInvalid     error = errors.New("这不是一个token,请重新登录")
)

// CreateToken  生成token
func (j *JWT) CreateToken(claims MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(j.JwtKey)
}

// ParseToken 解析token
func (j *JWT) ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.JwtKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	}
	return nil, TokenInvalid
}

// JwtToken Jwt中间件
func JwtToken(conf *viper.Viper) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		tokenHeader := ctx.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_EXIST
			ctx.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrorMsg(code),
			})
			ctx.Abort()
			return
		}
		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrorMsg(code),
			})
			ctx.Abort()
			return
		}

		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			ctx.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrorMsg(code),
			})
			ctx.Abort()
			return
		}

		j := NewJWT(conf)
		// 解析token
		claims, err := j.ParseToken(checkToken[1])
		if err != nil {
			if err == TokenExpired {
				ctx.JSON(http.StatusOK, gin.H{
					"status":  errmsg.ERROR,
					"message": "token授权已过期,请重新登录",
					"data":    nil,
				})
				ctx.Abort()
				return
			}
			ctx.JSON(http.StatusOK, gin.H{
				"status":  errmsg.ERROR,
				"message": err.Error(),
				"data":    nil,
			})
			ctx.Abort()
			return
		}
		ctx.Set("username", claims)
		ctx.Next()
	}
}
