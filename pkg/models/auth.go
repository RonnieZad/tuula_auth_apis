
package models

import (
    "gorm.io/gorm"
)

type User struct {
    Id       int64  `json:"id" gorm:"primaryKey"`
    PhoneNumber    string `json:"phoneNumber" gorm:"not null"`
    Name    string `json:"name" gorm:"not null;unique"`
    Email    string `json:"email" gorm:"not null"`
    Password string `json:"password" gorm:"not null"`
    gorm.Model
}