package models

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
	"time"
)

type User struct {
	gorm.Model
	Id        uint   `gorm:"primary_key"`
	FirstName string `gorm:"varchar(255);not null"`
	LastName  string `gorm:"varchar(255);not null"`
	Username  string `gorm:"column:username;unique_index"`
	Email     string `gorm:"column:email;unique_index"`
	Password  string `gorm:"column:password;not null"`

	Comments []Comment `gorm:"foreignkey:UserId"`

	Roles     []Role     `gorm:"many2many:users_roles;"`
	UserRoles []UserRole `gorm:"foreignkey:UserId"`
}

// SetPassword You can change the value in bcrypt.DefaultCost to adjust the security index.
// 	err := userModel.setPassword("password0")
func (u *User) SetPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty")
	}
	bytePassword := []byte(password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.Password = string(passwordHash)
	return nil
}

// IsValidPassword Database will only save the hashed string, you should check it by util function.
// 	if err := serModel.checkPassword("password0"); err != nil { password error }
func (u *User) IsValidPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func (u *User) BeforeSave(db *gorm.DB) (err error) {
	if len(u.Roles) == 0 {
		userRole := Role{}
		db.Model(&Role{}).Where("name = ?", "ROLE_USER").First(&userRole)
		u.Roles = append(u.Roles, userRole)
	}
	return
}

// GenerateJwtToken Generate JWT token associated to this user
func (u *User) GenerateJwtToken() string {
	jwtToken := jwt.New(jwt.SigningMethodHS512)

	var roles []string
	for _, role := range u.Roles {
		roles = append(roles, role.Name)
	}

	jwtToken.Claims = jwt.MapClaims{
		"user_id":  u.ID,
		"username": u.Username,
		"roles":    roles,
		"exp":      time.Now().Add(time.Hour * 24 * 90).Unix(),
	}
	// Sign and get the complete encoded token as a string
	token, _ := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return token
}

func (u *User) IsAdmin() bool {
	for _, role := range u.Roles {
		if role.Name == "ROLE_ADMIN" {
			return true
		}
	}
	return false
}
func (u *User) IsNotAdmin() bool {
	return !u.IsAdmin()
}
