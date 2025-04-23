package model

import (
	"time"

	"gorm.io/gorm"
)

// tambahan
type Mutasi struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey"`
	RekeningID uint   `gorm:"not null,index"`
	Tipe       string `gorm:"type:varchar(100)"` // "DEBIT" untuk tarik, "KREDIT" untuk tabung
	Nominal    int64  `gorm:"type:bigint"`
	Keterangan string `gorm:"type:varchar(255)"`
	CreatedAt  time.Time
	Rekening   Rekening `gorm:"foreignKey:RekeningID"`
}
