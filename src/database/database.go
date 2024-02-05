package database

import (
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"

  "api_pdv/src/controller/model/request"
)

var DB *gorm.DB

func SetupDB() error {
  db, err := gorm.Open(
    sqlite.Open("products.db"),
    &gorm.Config{},
  )
  if err != nil {
    return err
  }

  if err = db.AutoMigrate(&request.ProductRequest{}); err != nil {
    return err
  }

  DB = db

  return nil
}
