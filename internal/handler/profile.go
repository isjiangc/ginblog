package handler

import (
	"net/http"
	"strconv"

	"ginblog/internal/model"
	"ginblog/internal/service"
	"ginblog/pkg/helper/errmsg"
	"github.com/gin-gonic/gin"
)

type ProfileHandler interface {
	GetProfile(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
}

type profileHandler struct {
	*Handler
	profileService service.ProfileService
}

// GetProfile 获取个人信息设置
func (p profileHandler) GetProfile(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, code := p.profileService.GetProfile(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrorMsg(code),
	})
}

// UpdateProfile 更新个人信息设置
func (p profileHandler) UpdateProfile(ctx *gin.Context) {
	var data model.Profile
	id, _ := strconv.Atoi(ctx.Param("id"))
	_ = ctx.ShouldBindJSON(&data)

	code := p.profileService.UpdateProfile(id, &data)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMsg(code),
	})
}

func NewProfileHandler(handler *Handler, profileService service.ProfileService) ProfileHandler {
	return &profileHandler{
		Handler:        handler,
		profileService: profileService,
	}
}
