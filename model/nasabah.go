package model

import (
	"time"

	"gorm.io/gorm"
)

type Nasabah struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	NIK       string `gorm:"type:varchar(20);uniqueIndex"`
	Nama      string `gorm:"type:varchar(100)"`
	NoHP      string `gorm:"type:varchar(15)"`
	CreatedAt time.Time
}
