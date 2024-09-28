package dao

import "time"

type UserProfile struct {
	ID         int64     `gorm:"primaryKey;autoIncrement"`
	Name       string    `gorm:"type:varchar(32)"`
	Bio        string    `gorm:"type:varchar(100)"`
	Username   string    `gorm:"type:varchar(64)"`
	ImageURL   string    `gorm:"type:varchar(64)"`
	PostScript string    `gorm:"type:json"`      // Use `type:text` for MySQL if necessary
	CreatedAt  time.Time `gorm:"autoCreateTime"` // Automatically sets current timestamp on creation
	UpdatedAt  time.Time `gorm:"autoUpdateTime"` // Automatically updates timestamp on modification
}
