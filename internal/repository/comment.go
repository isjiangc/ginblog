package repository

import (
	"ginblog/internal/model"
	"ginblog/pkg/helper/errmsg"
	"gorm.io/gorm"
)

type CommentRepository interface {
	AddComment(data *model.Comment) int
	GetComment(id int) (model.Comment, int)
	GetCommentList(pageSize int, pageNum int) ([]model.Comment, int64, int)
	GetCommentCount(id int) int64
	GetCommentListFront(id int, pageSize int, pageNum int) ([]model.Comment, int64, int)
	DeleteComment(id uint) int
	CheckComment(id int, data *model.Comment) int
	UncheckComment(id int, data *model.Comment) int
}
type commentRepository struct {
	*Repository
}

// AddComment 新增评论
func (c commentRepository) AddComment(data *model.Comment) int {
	err := c.db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetComment 查询单个评论
func (c commentRepository) GetComment(id int) (model.Comment, int) {
	var comment model.Comment
	err := c.db.Where("id = ?", id).First(&comment).Error
	if err != nil {
		return comment, errmsg.ERROR
	}
	return comment, errmsg.SUCCESS
}

// GetCommentList 后台获取所有评论列表
func (c commentRepository) GetCommentList(pageSize int, pageNum int) ([]model.Comment, int64, int) {
	var commentList []model.Comment
	var total int64
	c.db.Find(&commentList).Count(&total)
	err := c.db.Model(&commentList).Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Select(
		"comment.id, article.title,user_id,article_id, user.username,comment.content, comment.status,comment.created_at,comment.deleted_at").Joins(
		"LEFT JOIN article ON comment.article_id = article.id").Joins(
		"LEFT JOIN user ON comment.user_id = user.id").Scan(&commentList).Error
	if err != nil {
		return commentList, 0, errmsg.ERROR
	}
	return commentList, total, errmsg.SUCCESS
}

// GetCommentCount 获取评论数量
func (c commentRepository) GetCommentCount(id int) int64 {
	var comment model.Comment
	var total int64
	c.db.Find(&comment).Where("article_id = ?", id).Where("status = ?", 1).Count(&total)
	return total
}

// GetCommentListFront 展示页码获取评论列表
func (c commentRepository) GetCommentListFront(id int, pageSize int, pageNum int) ([]model.Comment, int64, int) {
	var commentList []model.Comment
	var total int64
	c.db.Find(&model.Comment{}).Where("article_id = ?", id).Where("status = ?", 1).Count(&total)
	err := c.db.Model(&model.Comment{}).Limit(pageSize).Offset((pageNum-1)*pageSize).Order("Created_At DESC").Select(
		"comment.id , article.title, user_id, user.username, comment.content, comment.status, comment.created_at, comment.deleted_at").Joins(
		"LEFT JOIN article ON comment.article_id = article.id").Joins(
		"LEFT JOIN user ON comment.user_id = user.id").Where("article_id = ?", id).Where("status = ?", 1).Scan(&commentList).Error
	if err != nil {
		return commentList, 0, errmsg.ERROR
	}
	return commentList, total, errmsg.SUCCESS
}

// DeleteComment 删除评论
func (c commentRepository) DeleteComment(id uint) int {
	var comment model.Comment
	err := c.db.Where("id = ?", id).Delete(&comment).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// CheckComment 通过评论
func (c commentRepository) CheckComment(id int, data *model.Comment) int {
	var comment model.Comment
	var res model.Comment
	var article model.Article
	maps := make(map[string]interface{})
	maps["status"] = data.Status

	err := c.db.Model(&comment).Where("id = ?", id).Updates(maps).First(&res).Error
	c.db.Model(&article).Where("id = ?", res.ArticleId).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1))
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// UncheckComment 撤下评论
func (c commentRepository) UncheckComment(id int, data *model.Comment) int {
	var comment model.Comment
	var res model.Comment
	var article model.Article
	maps := make(map[string]interface{})
	maps["status"] = data.Status

	err := c.db.Model(&comment).Where("id = ?", id).Updates(maps).First(&res).Error
	c.db.Model(&article).Where("id = ?", res.ArticleId).UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1))
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func NewCommentRepository(repository *Repository) CommentRepository {
	return &commentRepository{
		Repository: repository,
	}
}
