package repository

import (
  "api_pdv/src/database"
  "api_pdv/src/controller/model/request"
)

func CreateProduct(product *request.ProductRequest) error {
  db := database.DB

  var last request.ProductRequest

  db.Last(&last)

  product.ID = last.ID + 1

  err := db.Create(product).Error

  return err
}
