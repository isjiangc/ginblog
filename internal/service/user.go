package service

import (
	"ginblog/internal/model"
	"ginblog/internal/repository"
)

type UserService interface {
	GetUserById(id int64) (*model.User, error)
	CheckUser(name string) (code int)
	CheckUpUser(id int, name string) (code int)
	CreateUser(data *model.User) int
	GetUser(id int) (model.User, int)
	GetUsers(username string, pageSize int, pageNum int) ([]model.User, int64)
	EditUser(id int, data *model.User) int
	ChangePassword(id int, data *model.User) int
	DeleteUser(id int) int
	CheckLogin(username string, password string) (model.User, int)
	CheckLoginFront(username string, password string) (model.User, int)
}

type userService struct {
	*Service
	userRepository repository.UserRepository
}

func (s *userService) CheckUser(name string) (code int) {
	return s.userRepository.CheckUser(name)
}

func (s *userService) CheckUpUser(id int, name string) (code int) {
	return s.userRepository.CheckUpUser(id, name)
}

func (s *userService) CreateUser(data *model.User) int {
	return s.userRepository.CreateUser(data)
}

func (s *userService) GetUser(id int) (model.User, int) {
	return s.userRepository.GetUser(id)
}

func (s *userService) GetUsers(username string, pageSize int, pageNum int) ([]model.User, int64) {
	return s.userRepository.GetUsers(username, pageSize, pageNum)
}

func (s *userService) EditUser(id int, data *model.User) int {
	return s.userRepository.EditUser(id, data)
}

func (s *userService) ChangePassword(id int, data *model.User) int {
	return s.userRepository.ChangePassword(id, data)
}

func (s *userService) DeleteUser(id int) int {
	return s.userRepository.DeleteUser(id)
}

func (s *userService) CheckLogin(username string, password string) (model.User, int) {
	return s.userRepository.CheckLogin(username, password)
}

func (s *userService) CheckLoginFront(username string, password string) (model.User, int) {
	return s.userRepository.CheckLoginFront(username, password)
}

func NewUserService(service *Service, userRepository repository.UserRepository) UserService {
	return &userService{
		Service:        service,
		userRepository: userRepository,
	}
}

func (s *userService) GetUserById(id int64) (*model.User, error) {
	return s.userRepository.FirstById(id)
}
