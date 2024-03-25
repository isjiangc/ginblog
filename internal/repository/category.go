package repository

import (
	"ginblog/internal/model"
	"ginblog/pkg/helper/errmsg"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	CheckCategory(name string) int
	CreateCate(data *model.Category) int
	GetCateInfo(id int) (model.Category, int)
	GetCate(pageSize int, pageNum int) ([]model.Category, int64)
	EditCate(id int, data *model.Category) int
	DeleteCate(id int) int
}

type categoryRepository struct {
	*Repository
}

// CheckCategory 查询分类是否存在
func (c categoryRepository) CheckCategory(name string) int {
	var cate model.Category
	c.db.Select("id").Where("name = ?", name).First(&cate)
	if cate.Id > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS
}

// CreateCate 新增分类
func (c categoryRepository) CreateCate(data *model.Category) int {
	err := c.db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetCateInfo 查询单个分类信息
func (c categoryRepository) GetCateInfo(id int) (model.Category, int) {
	var cate model.Category
	c.db.Where("id = ?", id).First(&cate)
	return cate, errmsg.SUCCESS
}

// GetCate 查询分类列表
func (c categoryRepository) GetCate(pageSize int, pageNum int) ([]model.Category, int64) {
	var cate []model.Category
	var total int64
	err := c.db.Find(&cate).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	c.db.Model(&cate).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cate, total
}

// EditCate 编辑分类信息
func (c categoryRepository) EditCate(id int, data *model.Category) int {
	var cate model.Category
	maps := make(map[string]interface{})
	maps["name"] = data.Name
	err := c.db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteCate 删除分类
func (c categoryRepository) DeleteCate(id int) int {
	var cate model.Category
	err := c.db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func NewCategoryRepository(repository *Repository) CategoryRepository {
	return &categoryRepository{
		Repository: repository,
	}
}
