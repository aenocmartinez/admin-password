package formrequest

type CreateSequenceFormRequest struct {
	Name   string `json:"name" binding:"required"`
	Prefix string `json:"prefix"`
	Value  int64  `json:"value"`
}
