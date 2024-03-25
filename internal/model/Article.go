package model

import "gorm.io/gorm"

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title        string `gorm:"column:title;type:varchar(100);NOT NULL" json:"title"`
	Cid          uint64 `gorm:"column:cid;type:bigint(20) unsigned;NOT NULL" json:"cid"`
	Desc         string `gorm:"column:desc;type:varchar(200)" json:"desc"`
	Content      string `gorm:"column:content;type:longtext" json:"content"`
	Img          string `gorm:"column:img;type:varchar(100)" json:"img"`
	CommentCount int64  `gorm:"column:comment_count;type:bigint(20);default:0;NOT NULL" json:"comment_count"`
	ReadCount    int64  `gorm:"column:read_count;type:bigint(20);default:0;NOT NULL" json:"read_count"`
}

func (m *Article) TableName() string {
	return "article"
}
