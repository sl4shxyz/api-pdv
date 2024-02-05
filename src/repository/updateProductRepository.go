package repository

import (
  "api_pdv/src/controller/model/request"
  "api_pdv/src/database"
)

func UpdateProduct(id string, productUpdate *request.ProductRequest) error {
  db := database.DB

  var product request.ProductRequest

  err := db.Find(&product, id).Error

  product.Name = productUpdate.Name
  product.Price = productUpdate.Price
  product.Amount = productUpdate.Amount

  err = db.Save(&product).Error

  return err
}
