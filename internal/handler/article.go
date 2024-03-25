package handler

import (
	"net/http"
	"strconv"

	"ginblog/internal/model"
	"ginblog/internal/service"
	"ginblog/pkg/helper/errmsg"
	"github.com/gin-gonic/gin"
)

type ArticleHandler interface {
	AddArticle(ctx *gin.Context)
	GetCateArt(ctx *gin.Context)
	GetArtInfo(ctx *gin.Context)
	GetArt(ctx *gin.Context)
	EditArt(ctx *gin.Context)
	DeleteArt(ctx *gin.Context)
}

type articleHandler struct {
	*Handler
	articleService service.ArticleService
}

func NewArticleHandler(handler *Handler, articleService service.ArticleService) ArticleHandler {
	return &articleHandler{
		Handler:        handler,
		articleService: articleService,
	}
}

// AddArticle 增加文章
func (a *articleHandler) AddArticle(ctx *gin.Context) {
	var data model.Article
	_ = ctx.ShouldBindJSON(&data)
	code := a.articleService.CreateArt(&data)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrorMsg(code),
	})
}

// GetCateArt 查询分类下的所有文章
func (a *articleHandler) GetCateArt(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	id, _ := strconv.Atoi(ctx.Param("id"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageNum <= 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	data, code, total := a.articleService.GetCateArt(id, pageSize, pageNum)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrorMsg(code),
	})
}

// GetArtInfo 查询单个文章信息
func (a *articleHandler) GetArtInfo(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, code := a.articleService.GetArtInfo(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrorMsg(code),
	})
}

// GetArt 查询文章列表
func (a *articleHandler) GetArt(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	title := ctx.Query("title")
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageNum <= 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	if len(title) == 0 {
		data, code, totle := a.articleService.GetArt(pageSize, pageNum)
		ctx.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   totle,
			"message": errmsg.GetErrorMsg(code),
		})
		return
	}
}

// EditArt 编辑文章
func (a *articleHandler) EditArt(ctx *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(ctx.Param("id"))
	_ = ctx.ShouldBindJSON(&data)

	code := a.articleService.EditArt(id, &data)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMsg(code),
	})
}

// DeleteArt 删除文章
func (a *articleHandler) DeleteArt(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	code := a.articleService.DeleteArt(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrorMsg(code),
	})
}
