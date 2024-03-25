package handler

import (
	"net/http"
	"strconv"

	"ginblog/internal/model"
	"ginblog/internal/service"
	"ginblog/pkg/helper/errmsg"
	"ginblog/pkg/helper/resp"
	"ginblog/pkg/helper/validator"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler interface {
	GetUserById(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	AddUser(ctx *gin.Context)
	GetUserInfo(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
	EditUser(ctx *gin.Context)
	ChangeUserPassword(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userHandler struct {
	*Handler
	userService service.UserService
}

// AddUser 添加用户
func (h *userHandler) AddUser(ctx *gin.Context) {
	var data model.User
	var msg string
	var validCode int
	_ = ctx.ShouldBindJSON(&data)
	msg, validCode = validator.Validate(&data)
	if validCode != errmsg.SUCCESS {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  validCode,
			"message": msg,
		})
		ctx.Abort()
		return
	}
	code := h.userService.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		h.userService.CreateUser(&data)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMsg(code),
	})
}

// GetUserInfo 查询单个用户
func (h *userHandler) GetUserInfo(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	maps := make(map[string]interface{})
	data, code := h.userService.GetUser(id)
	maps["username"] = data.Username
	maps["role"] = data.Role
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    maps,
		"total":   1,
		"message": errmsg.GetErrorMsg(code),
	})
}

// GetUsers 查询用户列表
func (h *userHandler) GetUsers(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	username := ctx.Query("username")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	data, total := h.userService.GetUsers(username, pageSize, pageNum)
	code := errmsg.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrorMsg(code),
	})
}

// EditUser 编辑用户
func (h *userHandler) EditUser(ctx *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(ctx.Param("id"))
	_ = ctx.ShouldBindJSON(&data)

	code := h.userService.CheckUpUser(id, data.Username)
	if code == errmsg.SUCCESS {
		h.userService.EditUser(id, &data)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMsg(code),
	})
}

// ChangeUserPassword 修改密码
func (h *userHandler) ChangeUserPassword(ctx *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(ctx.Param("id"))
	_ = ctx.ShouldBindJSON(&data)

	code := h.userService.ChangePassword(id, &data)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMsg(code),
	})
}

// DeleteUser 删除用户
func (h *userHandler) DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	code := h.userService.DeleteUser(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMsg(code),
	})
}

func NewUserHandler(handler *Handler, userService service.UserService) UserHandler {
	return &userHandler{
		Handler:     handler,
		userService: userService,
	}
}

func (h *userHandler) GetUserById(ctx *gin.Context) {
	var params struct {
		Id int64 `form:"id" binding:"required"`
	}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	user, err := h.userService.GetUserById(params.Id)
	h.logger.Info("GetUserByID", zap.Any("user", user))
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	resp.HandleSuccess(ctx, user)
}

func (h *userHandler) UpdateUser(ctx *gin.Context) {
	resp.HandleSuccess(ctx, nil)
}
