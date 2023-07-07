package dto

type RegisterReq struct {
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
