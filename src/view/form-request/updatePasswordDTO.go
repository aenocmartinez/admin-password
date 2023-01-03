package formrequest

type UpdatePasswordDTOFormRequest struct {
	Id       int64  `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
	Url      string `json:"url" binding:"required"`
}
