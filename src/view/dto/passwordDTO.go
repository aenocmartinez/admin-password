package dto

type PasswordDTO struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Url      string `json:"url"`
}
