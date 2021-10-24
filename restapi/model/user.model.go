package model

// User Model
type User struct {
	ID        uint64 `json:"_id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	UUID      string `json:"uuid"`
	Confirmed bool   `json:"confirmed"`
}

// ConfirmData ---  Used in ConfirmAccount handler
type ConfirmData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UUID     string `json:"uuid"`
}

// LoginData --- Used in Login handler
type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RequestPasswordData -- Used in RequestPassword
type RequestPasswordData struct {
	Email string `json:"email"`
}

// ForgotPasswordData --- Used in ForgotPassword handler
type ForgotPasswordData struct {
	Email       string `json:"email"`
	UUID        string `json:"uuid"`
	NewPassword string `json:"newPassword"`
}
