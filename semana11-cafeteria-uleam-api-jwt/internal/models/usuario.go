package models

import "time"

type Usuario struct {
    ID           int
    Email        string    `json:"email" gorm:"not null;uniqueIndex"`
    PasswordHash string    `json:"-" gorm:"not null"`
    CreatedAt    time.Time `json:"created_at"`
}
