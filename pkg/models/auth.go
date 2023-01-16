package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID                          uuid.UUID `json:"id" gorm:"primary_key"`
	PhoneNumber                 string    `json:"phoneNumber"`
	PhoneNumberVerificationCode string    `json:"phoneNumberVerificationCode"`
	Name                        string    `json:"name"`
	EmailAddress                string    `json:"emailAddress"`
	DateOfBirth                 string    `json:"dateOfBirth"`
	Password                    string    `json:"password"`
	IsKYCVerified               bool      `json:"kycVerified"`
	CreditScore                 float32   `json:"creditScore"`
	IsFinanceWorthy             bool      `json:"isFinanceWorthy"`
	WorkPlace                   string    `json:"workPlace"`
	NIN                         string    `json:"nin"`
	EmployerName                string    `json:"employerName"`
	SalaryScale                 float32   `json:"salaryScale"`
	gorm.Model                  `json:"-"`
}
