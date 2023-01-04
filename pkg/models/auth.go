
package models

type User struct {
    Id       int64  `json:"id" gorm:"primaryKey"`
    PhoneNumber    string `json:"phoneNumber"`
    Name    string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}