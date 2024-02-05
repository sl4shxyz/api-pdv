package request

type SaleRequest struct {
  Amount int `json:"amount" binding:"required,min=1"`
}
