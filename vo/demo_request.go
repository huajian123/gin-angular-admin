package vo

type CreateDemoRequest struct {
	Name string `json:"name" binding:"required"`
}
