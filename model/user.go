package model

import "gorm.io/gorm"

// User ...
type User struct {
	gorm.Model
	Name string
}

// Users ...
type Users []*User

// UpdateUser ...
func (u *User) UpdateUser(reqUser *User) {
	if reqUser.Name != "" {
		u.Name = reqUser.Name
	}
}
