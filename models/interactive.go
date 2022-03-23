package models

import "gorm.io/gorm"

type Interactive struct {
	gorm.Model

	Uid uint `gorm:"not null"`

	Vid   uint  `gorm:"not null"`
	Video Video `gorm:"ForeignKey:id;AssociationForeignKey:vid"`

	Favorite bool `gorm:"default:false"` //是否收藏
	//like和SQL的关键词冲突了，查询时需要写成`like`
	Like bool `gorm:"default:false"` //是否点赞
}

// 还未将 favorite 功能继承进入，但 like 功能依赖于此结构体
