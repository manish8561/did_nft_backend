package clients

type ClientValidator struct {
	Email string `form:"email" json:"email" binding:"required,email,max=100"`
}

type ClientStatusValidator struct {
	ID     uint   `form:"id" json:"id" binding:"required,number"`
	Status string `form:"status" json:"status" binding:"required,max=50"`
}
