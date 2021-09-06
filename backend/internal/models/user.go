package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	//Id           uint    `gorm:"primary_key"`
	FirstName string `gorm:"varchar(255);not null"`
	LastName  string `gorm:"varchar(255);not null"`
	Username  string `gorm:"column:username"`
	Email     string `gorm:"column:email;unique_index"`
	Password  string `gorm:"column:password;not null"`

	Comments []Comment `gorm:"foreignkey:UserId"`

	Roles     []Role     `gorm:"many2many:users_roles;"`
	UserRoles []UserRole `gorm:"foreignkey:UserId"`
}
