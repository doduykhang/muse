package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseID struct {
	ID uint
}

type BaseModel struct {
	ID uint `gorm:"primarykey"`
	Audit
}

type Audit struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	CreatedBy string
	UpdatedBy string
	DeletedBy string
}

type AuditHardDelete struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy string
	UpdatedBy string
}
