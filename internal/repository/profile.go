package repository

import (
	"ginblog/internal/model"
	"ginblog/pkg/helper/errmsg"
)

type ProfileRepository interface {
	GetProfile(id int) (model.Profile, int)
	UpdateProfile(id int, data *model.Profile) int
}

type profileRepository struct {
	*Repository
}

// GetProfile 获取个人信息设置
func (p profileRepository) GetProfile(id int) (model.Profile, int) {
	var profile model.Profile
	err := p.db.Where("ID = ?", id).First(&profile).Error
	if err != nil {
		return profile, errmsg.ERROR
	}
	return profile, errmsg.SUCCESS
}

// UpdateProfile 更新个人信息设置
func (p profileRepository) UpdateProfile(id int, data *model.Profile) int {
	var profile model.Profile
	err := p.db.Model(&profile).Where("ID = ?", id).Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func NewProfileRepository(repository *Repository) ProfileRepository {
	return &profileRepository{
		Repository: repository,
	}
}
