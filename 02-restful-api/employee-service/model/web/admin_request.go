package web

type AdminLoginRequest struct {
	Username string `validate:"required,max=30" json:"username"`
	Password string `validate:"required,max=30" json:"password"`
}
