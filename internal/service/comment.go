package service

import (
	"ginblog/internal/model"
	"ginblog/internal/repository"
)

type CommentService interface {
	AddComment(data *model.Comment) int
	GetComment(id int) (model.Comment, int)
	GetCommentList(pageSize int, pageNum int) ([]model.Comment, int64, int)
	GetCommentCount(id int) int64
	GetCommentListFront(id int, pageSize int, pageNum int) ([]model.Comment, int64, int)
	DeleteComment(id uint) int
	CheckComment(id int, data *model.Comment) int
	UncheckComment(id int, data *model.Comment) int
}

type commentService struct {
	*Service
	commentRepository repository.CommentRepository
}

// AddComment 新增评论
func (c commentService) AddComment(data *model.Comment) int {
	return c.commentRepository.AddComment(data)
}

// GetComment 查询单个评论
func (c commentService) GetComment(id int) (model.Comment, int) {
	return c.commentRepository.GetComment(id)
}

// GetCommentList 后台获取所有评论列表
func (c commentService) GetCommentList(pageSize int, pageNum int) ([]model.Comment, int64, int) {
	return c.commentRepository.GetCommentList(pageSize, pageNum)
}

// GetCommentCount 获取评论数量
func (c commentService) GetCommentCount(id int) int64 {
	return c.commentRepository.GetCommentCount(id)
}

// GetCommentListFront 展示页码获取评论列表
func (c commentService) GetCommentListFront(id int, pageSize int, pageNum int) ([]model.Comment, int64, int) {
	return c.commentRepository.GetCommentListFront(id, pageSize, pageNum)
}

// DeleteComment 删除评论
func (c commentService) DeleteComment(id uint) int {
	return c.commentRepository.DeleteComment(id)
}

// CheckComment 通过评论
func (c commentService) CheckComment(id int, data *model.Comment) int {
	return c.commentRepository.CheckComment(id, data)
}

// UncheckComment 撤下评论
func (c commentService) UncheckComment(id int, data *model.Comment) int {
	return c.commentRepository.UncheckComment(id, data)
}

func NewCommentService(service *Service, commentRepository repository.CommentRepository) CommentService {
	return &commentService{
		Service:           service,
		commentRepository: commentRepository,
	}
}
