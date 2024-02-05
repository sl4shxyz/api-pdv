package request

type ProductRequest struct {
  ID     int      `json:"id"      gorm:"primaryKey;autoIncrement"`
  Name   string   `json:"name"    binding:"required,max=500"`
  Price  float32  `json:"price"   binding:"required"`
  Amount int      `json:"amount"  binding:"required"`
}
