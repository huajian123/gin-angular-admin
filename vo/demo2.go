package vo

// 承接参数
type CreateDemo2Request struct {
	DemoId  uint   `json:"demoId" binding:"required"` // bingding表单验证
	Title   string `json:"title" binding:"required,max=10"`
	HeadImg string `json:"head_img"`
	Content string `json:"content" binding:"required"`
}
