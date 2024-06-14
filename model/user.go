package model

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ChangePasswordRequest struct {
	Email              string `json:"email" db:"email"`
	PasswordSekarang   string `json:"password_sekarang"`
	PasswordBaru       string `json:"password_baru"`
	PasswordBaruRepeat string `json:"password_baru_repeat"`
}

type User struct {
	Email    string `json:"email" db:"email" validate:"required,email"`
	Password string `json:"password" db:"password" validate:"required"`
	IsActive string `json:"is_active" db:"is_active"`
	Unit     string `json:"unit" db:"unit"`
}

type UserResponse struct {
	Email    string `json:"email" db:"email"`
	IsActive string `json:"is_active" db:"is_active"`
	Unit     string `json:"unit" db:"unit"`
}
