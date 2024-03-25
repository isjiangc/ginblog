package service

import (
	"ginblog/internal/model"
	"ginblog/internal/repository"
)

type ArticleService interface {
	CreateArt(data *model.Article) int
	GetCateArt(id int, pageSize int, pageNum int) ([]model.Article, int, int64)
	GetArtInfo(id int) (model.Article, int)
	GetArt(pageSize int, pageNum int) ([]model.Article, int, int64)
	EditArt(id int, data *model.Article) int
	DeleteArt(id int) int
}
type articleService struct {
	*Service
	articleRepository repository.ArticleRepository
}

// CreateArt 新增文章
func (a articleService) CreateArt(data *model.Article) int {
	return a.articleRepository.CreateArt(data)
}

// GetCateArt 查询分类下的所有文章
func (a articleService) GetCateArt(id int, pageSize int, pageNum int) ([]model.Article, int, int64) {
	return a.articleRepository.GetCateArt(id, pageSize, pageNum)
}

// GetArtInfo 查询单个文章
func (a articleService) GetArtInfo(id int) (model.Article, int) {
	return a.articleRepository.GetArtInfo(id)
}

// GetArt 查询文章列表
func (a articleService) GetArt(pageSize int, pageNum int) ([]model.Article, int, int64) {
	return a.articleRepository.GetArt(pageSize, pageNum)
}

// EditArt 编辑文章
func (a articleService) EditArt(id int, data *model.Article) int {
	return a.articleRepository.EditArt(id, data)
}

// DeleteArt 删除文章
func (a articleService) DeleteArt(id int) int {
	return a.articleRepository.DeleteArt(id)
}

func NewArticleService(service *Service, articleRepository repository.ArticleRepository) ArticleService {
	return &articleService{
		Service:           service,
		articleRepository: articleRepository,
	}
}
