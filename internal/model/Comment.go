package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserId       uint64 `gorm:"column:user_id;type:bigint(20) unsigned" json:"user_id"`
	ArticleId    uint64 `gorm:"column:article_id;type:bigint(20) unsigned" json:"article_id"`
	Content      string `gorm:"column:content;type:varchar(500);NOT NULL" json:"content"`
	Status       int    `gorm:"column:status;type:tinyint(4);default:2" json:"status"`
	ArticleTitle string `gorm:"column:article_title;type:longtext" json:"article_title"`
	Username     string `gorm:"column:username;type:longtext" json:"username"`
	Title        string `gorm:"column:title;type:longtext" json:"title"`
}

func (m *Comment) TableName() string {
	return "comment"
}
