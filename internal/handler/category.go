package handler

import (
	"net/http"
	"strconv"

	"ginblog/internal/model"
	"ginblog/internal/service"
	"ginblog/pkg/helper/errmsg"
	"github.com/gin-gonic/gin"
)

type CategoryHandler interface {
	AddCategory(ctx *gin.Context)
	GetCateInfo(ctx *gin.Context)
	GetCate(ctx *gin.Context)
	EditCate(ctx *gin.Context)
	DeleteCate(ctx *gin.Context)
}
type categoryHandler struct {
	*Handler
	categoryService service.CategoryService
}

// AddCategory 添加分类
func (c categoryHandler) AddCategory(ctx *gin.Context) {
	var data model.Category
	_ = ctx.ShouldBindJSON(&data)
	code := c.categoryService.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		c.categoryService.CreateCate(&data)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrorMsg(code),
	})
}

// GetCateInfo 查询分类信息
func (c categoryHandler) GetCateInfo(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, code := c.categoryService.GetCateInfo(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrorMsg(code),
	})
}

// GetCate 查询分类列表
func (c categoryHandler) GetCate(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageNum <= 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	data, total := c.categoryService.GetCate(pageSize, pageNum)
	code := errmsg.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrorMsg(code),
	})
}

// EditCate 编辑分类名
func (c categoryHandler) EditCate(ctx *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(ctx.Param("id"))
	_ = ctx.ShouldBindJSON(&data)
	code := c.categoryService.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		c.categoryService.EditCate(id, &data)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		ctx.Abort()
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMsg(code),
	})
}

// DeleteCate 删除分类
func (c categoryHandler) DeleteCate(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	code := c.categoryService.DeleteCate(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMsg(code),
	})
}

func NewCategoryHandler(handler *Handler, categoryService service.CategoryService) CategoryHandler {
	return &categoryHandler{
		Handler:         handler,
		categoryService: categoryService,
	}
}
