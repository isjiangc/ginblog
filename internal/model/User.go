package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"column:username;type:varchar(20);NOT NULL" json:"username"`
	Password string `gorm:"column:password;type:varchar(500);NOT NULL" json:"password"`
	Role     int64  `gorm:"column:role;type:bigint(20);default:2" json:"role"`
}

func (m *User) TableName() string {
	return "user"
}
