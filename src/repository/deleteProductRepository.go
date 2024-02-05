package repository

import (
  "api_pdv/src/database"
  "api_pdv/src/controller/model/request"
)

func DeleteProduct(id string) error {
  db := database.DB

  var product request.ProductRequest

  err := db.Delete(&product, id).Error

  return err
}
