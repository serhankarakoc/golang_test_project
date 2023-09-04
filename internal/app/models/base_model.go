package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	CreatedBy uint
	UpdatedAt time.Time
	UpdatedBy uint
	DeletedAt gorm.DeletedAt `gorm:"index"`
	DeletedBy uint
}

func (bm *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	currentTime := time.Now()
	if bm.CreatedAt.IsZero() {
		bm.CreatedAt = currentTime
	}
	if bm.UpdatedAt.IsZero() {
		bm.UpdatedAt = currentTime
	}
	if bm.CreatedBy == 0 {
		bm.CreatedBy = 1
	}
	return
}

func (bm *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	bm.UpdatedAt = time.Now()
	if bm.UpdatedBy == 0 {
		bm.UpdatedBy = 1
	}
	return
}

func (bm *BaseModel) BeforeDelete(tx *gorm.DB) (err error) {
	bm.DeletedBy = 1
	return
}
