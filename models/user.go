package models

import "time"

// User represents the users table.
type User struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Email       string    `gorm:"type:varchar(255);not null;uniqueIndex" json:"email"`
	PhoneNumber string    `gorm:"type:varchar(50)" json:"phone_number"`
	Password    string    `gorm:"type:text;not null" json:"-"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"not null" json:"updated_at"`
}
