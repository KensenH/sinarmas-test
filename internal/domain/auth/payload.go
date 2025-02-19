package auth

type LoginReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"password"`
}
