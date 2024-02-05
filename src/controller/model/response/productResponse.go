package response

type ProductResponse struct {
  ID     int     `json:"id"`
  Name   string  `json:"name"`
  Price  float32 `json:"price"`
  Amount int     `json:"amount"`
}
