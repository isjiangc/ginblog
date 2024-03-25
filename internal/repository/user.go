package repository

import (
	"ginblog/internal/model"
	"ginblog/pkg/helper/errmsg"
	"ginblog/pkg/helper/scrypt"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	FirstById(id int64) (*model.User, error)
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
type userRepository struct {
	*Repository
}

// CheckUser 查询用户是否存在
func (r *userRepository) CheckUser(name string) (code int) {
	var user model.User
	r.db.Select("id").Where("username = ?", name).First(&user)

	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// CheckUpUser 更新查询
func (r *userRepository) CheckUpUser(id int, name string) (code int) {
	var user model.User
	r.db.Select("id, username").Where("username = ?", name).First(&user)
	if user.ID == uint(id) {
		return errmsg.SUCCESS
	}
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// CreateUser 新增用户
func (r *userRepository) CreateUser(data *model.User) int {
	data.Password = scrypt.ScryptPw(data.Password)
	err := r.db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetUser 查询用户
func (r *userRepository) GetUser(id int) (model.User, int) {
	var user model.User
	err := r.db.Limit(1).Where("ID = ?", id).Find(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCESS
}

// GetUsers 查询用户列表
func (r *userRepository) GetUsers(username string, pageSize int, pageNum int) ([]model.User, int64) {
	var users []model.User
	var total int64

	if username != "" {
		r.db.Select("id, username, role, created_at").Where(
			"username LIKE ?", username+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
		r.db.Model(&users).Where(
			"username LIKE ?", username+"%").Count(&total)
		return users, total
	}
	r.db.Select("id, username, role, created_at").Limit(pageSize).Offset((pageNum - 1) * pageNum).Find(&users)
	err := r.db.Model(&users).Count(&total)
	if err != nil {
		return users, 0
	}
	return users, total
}

// EditUser 编辑用户信息
func (r *userRepository) EditUser(id int, data *model.User) int {
	var user model.User
	maps := make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := r.db.Model(&user).Where("id  = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// ChangePassword 修改密码
func (r *userRepository) ChangePassword(id int, data *model.User) int {
	var user model.User
	maps := make(map[string]interface{})
	maps["password"] = scrypt.ScryptPw(data.Password)
	err := r.db.Model(&user).Where("id = ?", id).Updates(maps).Error
	// err := r.db.Select("password").Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteUser 删除用户
func (r *userRepository) DeleteUser(id int) int {
	var user model.User
	err := r.db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// CheckLogin 后台登录验证
func (r *userRepository) CheckLogin(username string, password string) (model.User, int) {
	var user model.User
	var PasswordErr error
	r.db.Where("username = ?", username).First(&user)
	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	if PasswordErr != nil {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return user, errmsg.ERROR_USER_NOT_RIGHT
	}
	return user, errmsg.SUCCESS
}

// CheckLoginFront 前台登录
func (r *userRepository) CheckLoginFront(username string, password string) (model.User, int) {
	var user model.User
	var PasswordErr error
	r.db.Where("username = ?", username).First(&user)
	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	if PasswordErr != nil {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	return user, errmsg.SUCCESS
}

func NewUserRepository(repository *Repository) UserRepository {
	return &userRepository{
		Repository: repository,
	}
}

func (r *userRepository) FirstById(id int64) (*model.User, error) {
	var user model.User
	// TODO: query db
	return &user, nil
}
