package dto

type RegisterRes struct {
	ID         int64  `json:"id"`
	FullName   string `json:"full_name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	RefrenceID string `json:"refrence_id"`
}
