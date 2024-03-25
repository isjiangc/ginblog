package handler

import (
	"net/http"
	"strconv"

	"ginblog/internal/model"
	"ginblog/internal/service"
	"ginblog/pkg/helper/errmsg"
	"github.com/gin-gonic/gin"
)

type CommentHandler interface {
	AddComment(ctx *gin.Context)
	GetComment(ctx *gin.Context)
	DeleteComment(ctx *gin.Context)
	GetCommentCont(ctx *gin.Context)
	GetCommentList(ctx *gin.Context)
	GetCommentListFront(ctx *gin.Context)
	CheckComment(ctx *gin.Context)
	UncheckComment(ctx *gin.Context)
}

type commentHandler struct {
	*Handler
	commentService service.CommentService
}

// AddComment 新增评论
func (c commentHandler) AddComment(ctx *gin.Context) {
	var data model.Comment
	_ = ctx.ShouldBindJSON(&data)

	code := c.commentService.AddComment(&data)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrorMsg(code),
	})
}

// GetComment 获取单个评论信息
func (c commentHandler) GetComment(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, code := c.commentService.GetComment(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrorMsg(code),
	})
}

// DeleteComment 删除评论
func (c commentHandler) DeleteComment(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	code := c.commentService.DeleteComment(uint(id))
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMsg(code),
	})
}

// GetCommentCont 获取评论数量
func (c commentHandler) GetCommentCont(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	total := c.commentService.GetCommentCount(id)
	ctx.JSON(http.StatusOK, gin.H{
		"total": total,
	})
}

// GetCommentList 后台查询评论列表
func (c commentHandler) GetCommentList(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	data, total, code := c.commentService.GetCommentList(pageSize, pageNum)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrorMsg(code),
	})
}

// GetCommentListFront 展示页面显示评论列表
func (c commentHandler) GetCommentListFront(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	data, total, code := c.commentService.GetCommentListFront(id, pageSize, pageNum)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrorMsg(code),
	})
}

// CheckComment 通过审核
func (c commentHandler) CheckComment(ctx *gin.Context) {
	var data model.Comment
	_ = ctx.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(ctx.Param("id"))

	code := c.commentService.CheckComment(id, &data)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMsg(code),
	})
}

// UncheckComment 撤下评论审核
func (c commentHandler) UncheckComment(ctx *gin.Context) {
	var data model.Comment
	_ = ctx.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(ctx.Param("id"))

	code := c.commentService.UncheckComment(id, &data)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMsg(code),
	})
}

func NewCommentHandler(handler *Handler, commentService service.CommentService) CommentHandler {
	return &commentHandler{
		Handler:        handler,
		commentService: commentService,
	}
}
