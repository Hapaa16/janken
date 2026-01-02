package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Email     string    `gorm:"uniqueIndex;not null"`
	Username  string    `gorm:"uniqueIndex;not null"`
	Avatar    string    `gorm:"not null"`
	Rank      int32     `gorm:"not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:now()"`
	UpdatedAt time.Time `gorm:"default:now()"`
}
