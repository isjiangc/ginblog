package service

import (
	"ginblog/internal/model"
	"ginblog/internal/repository"
)

type CategoryService interface {
	CheckCategory(name string) int
	CreateCate(data *model.Category) int
	GetCateInfo(id int) (model.Category, int)
	GetCate(pageSize int, pageNum int) ([]model.Category, int64)
	EditCate(id int, data *model.Category) int
	DeleteCate(id int) int
}

type categoryService struct {
	*Service
	categoryRepository repository.CategoryRepository
}

// CheckCategory 查询分类是否存在
func (c categoryService) CheckCategory(name string) int {
	return c.categoryRepository.CheckCategory(name)
}

// CreateCate 新增分类
func (c categoryService) CreateCate(data *model.Category) int {
	return c.categoryRepository.CreateCate(data)
}

// GetCateInfo 查询单个分类信息
func (c categoryService) GetCateInfo(id int) (model.Category, int) {
	return c.categoryRepository.GetCateInfo(id)
}

// GetCate 查询分类列表
func (c categoryService) GetCate(pageSize int, pageNum int) ([]model.Category, int64) {
	return c.categoryRepository.GetCate(pageSize, pageNum)
}

// EditCate 编辑分类信息
func (c categoryService) EditCate(id int, data *model.Category) int {
	return c.categoryRepository.EditCate(id, data)
}

// DeleteCate 删除分类
func (c categoryService) DeleteCate(id int) int {
	return c.categoryRepository.DeleteCate(id)
}

func NewCategoryService(service *Service, categoryRepository repository.CategoryRepository) CategoryService {
	return &categoryService{
		Service:            service,
		categoryRepository: categoryRepository,
	}
}
