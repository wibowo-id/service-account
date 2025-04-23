package model

import (
	"time"

	"gorm.io/gorm"
)

type Rekening struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey"`
	NoRekening string `gorm:"type:varchar(20);uniqueIndex"`
	NasabahID  uint   `gorm:"not null,index"`
	Saldo      int64  `gorm:"type:bigint"`
	CreatedAt  time.Time
	Nasabah    Nasabah `gorm:"foreignKey:NasabahID"`
}
