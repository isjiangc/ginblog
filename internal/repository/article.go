package repository

import (
	"ginblog/internal/model"
	"ginblog/pkg/helper/errmsg"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	CreateArt(data *model.Article) int
	GetCateArt(id int, pageSize int, pageNum int) ([]model.Article, int, int64)
	GetArtInfo(id int) (model.Article, int)
	GetArt(pageSize int, pageNum int) ([]model.Article, int, int64)
	EditArt(id int, data *model.Article) int
	DeleteArt(id int) int
}

type articleRepository struct {
	*Repository
}

func NewArticleRepository(repository *Repository) ArticleRepository {
	return &articleRepository{
		Repository: repository,
	}
}

// CreateArt 增加文章
func (r *articleRepository) CreateArt(data *model.Article) int {
	err := r.db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetCateArt 查询分类下的所有文章
func (r *articleRepository) GetCateArt(id int, pageSize int, pageNum int) ([]model.Article, int, int64) {
	var cateArtList []model.Article
	var total int64
	err := r.db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where(
		"cid =?", id).Find(&cateArtList).Error
	r.db.Model(&cateArtList).Where("cid =?", id).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return cateArtList, errmsg.SUCCESS, total
}

// GetArtInfo 查询单个文章
func (r *articleRepository) GetArtInfo(id int) (model.Article, int) {
	var art model.Article
	err := r.db.Where("id = ?", id).Preload("Category").First(&art).Error
	r.db.Model(&art).Where("id = ?", id).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1))
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCESS
}

// GetArt 查询文章列表
func (r *articleRepository) GetArt(pageSize int, pageNum int) ([]model.Article, int, int64) {
	var articleList []model.Article
	var err error
	var total int64
	err = r.db.Select("article.id, title, img, created_at, updated_at, `desc`, comment_count, read_count, category.name").Limit(
		pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Joins("Category").Find(&articleList).Error
	r.db.Model(&articleList).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCESS, total
}

// EditArt 编辑文章
func (r *articleRepository) EditArt(id int, data *model.Article) int {
	var art model.Article
	maps := make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err := r.db.Model(&art).Where("id = ?", id).Updates(&maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteArt 删除文章
func (r *articleRepository) DeleteArt(id int) int {
	var art model.Article
	err := r.db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
