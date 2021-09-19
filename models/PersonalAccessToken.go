package models

import (
	"database/sql"
	"time"
)

type PersonalAccessToken struct {
	ID         uint64 `gorm:"autoIncrement,primaryKey"`
	UserId     uint64
	Name       string   `gorm:"uniqueIndex,size:255"`
	Token      string   `gorm:"uniqueIndex,size:64"`
	Abilities  []string `gorm:"type:json"`
	LastUsedAt sql.NullTime
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoCreateTime,autoUpdateTime"`
}
