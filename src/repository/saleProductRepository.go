package repository

import (
  "api_pdv/src/database"
  "api_pdv/src/controller/model/request"

  "errors"
)

func SaleProduct(id string, sale *request.SaleRequest) error {
  db := database.DB

  var product request.ProductRequest

  if err := db.Find(&product, id).Error; err != nil {
    return err
  }

  if sale.Amount > product.Amount || sale.Amount < 0 {
    return errors.New("Invalid product quantity")
  }

  product.Amount -= sale.Amount

  err := db.Save(&product).Error

  return err
}
