package formrequest

type CreatePasswordFormRequest struct {
	Name     string `json:"name" binding:"required"`
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
	Url      string `json:"url" binding:"required"`
}
