package model

type Profile struct {
	Id     int64  `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	Name   string `gorm:"column:name;type:varchar(20)" json:"name"`
	Desc   string `gorm:"column:desc;type:varchar(200)" json:"desc"`
	Qqchat string `gorm:"column:qqchat;type:varchar(200)" json:"qqchat"`
	Wechat string `gorm:"column:wechat;type:varchar(100)" json:"wechat"`
	Weibo  string `gorm:"column:weibo;type:varchar(200)" json:"weibo"`
	Bili   string `gorm:"column:bili;type:varchar(200)" json:"bili"`
	Email  string `gorm:"column:email;type:varchar(200)" json:"email"`
	Img    string `gorm:"column:img;type:varchar(200)" json:"img"`
	Avatar string `gorm:"column:avatar;type:varchar(200)" json:"avatar"`
}

func (m *Profile) TableName() string {
	return "profile"
}
