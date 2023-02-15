package entity

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name       string `valid : -`
	Email      string `valid :"email"`
	EmployeeID string `valid :"matches(^[M]//(8)$)"`
}
