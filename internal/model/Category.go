package model

type Category struct {
	Id   uint64 `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Name string `gorm:"column:name;type:varchar(20);NOT NULL" json:"name"`
}

func (m *Category) TableName() string {
	return "category"
}
