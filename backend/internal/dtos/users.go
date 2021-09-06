package dtos

type RegisterRequestDto struct {
	Username             string `form:"username" json:"username" xml:"username"  binding:"required"`
	FirstName            string `form:"first_name" json:"first_name" xml:"first_name" binding:"required"`
	LastName             string `form:"last_name" json:"last_name" xml:"last_name" binding:"required"`
	Email                string `form:"email" json:"email" xml:"email" binding:"required"`
	Password             string `form:"password" json:"password" xml:"password" binding:"required"`
	PasswordConfirmation string `form:"password_confirmation" json:"password_confirmation" xml:"password-confirmation" binding:"required"`
}

type LoginRequestDto struct {
	// Username string `form:"username" json:"username" xml:"username" binding:"exists,username"`
	Username string `form:"username" json:"username" xml:"username" binding:"required"`
	Password string `form:"password"json:"password" binding:"exists,min=8,max=255"`

	userModel models.User `json:"-"`
}
