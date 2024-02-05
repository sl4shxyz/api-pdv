package repository

import (
  "api_pdv/src/database"
  "api_pdv/src/controller/model/request"
)

func GetProducts() ([]request.ProductRequest, error) {
  db := database.DB

  var products []request.ProductRequest

  err := db.Find(&products).Error

  return products, err
}

func GetProduct(id string) (request.ProductRequest, error) {
  db := database.DB

  var product request.ProductRequest

  err := db.First(&product, id).Error

  return product, err
}
