package service

import (
	"ginblog/internal/model"
	"ginblog/internal/repository"
)

type ProfileService interface {
	GetProfile(id int) (model.Profile, int)
	UpdateProfile(id int, data *model.Profile) int
}
type profileService struct {
	*Service
	profileRepository repository.ProfileRepository
}

// GetProfile 获取个人信息设置
func (p profileService) GetProfile(id int) (model.Profile, int) {
	return p.profileRepository.GetProfile(id)
}

// UpdateProfile 更新个人信息设置
func (p profileService) UpdateProfile(id int, data *model.Profile) int {
	return p.profileRepository.UpdateProfile(id, data)
}

func NewProfileService(service *Service, profileRepository repository.ProfileRepository) ProfileService {
	return &profileService{
		Service:           service,
		profileRepository: profileRepository,
	}
}
