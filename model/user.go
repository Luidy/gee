package model

import "gorm.io/gorm"

// User ...
type User struct {
	gorm.Model
	Name string
}
