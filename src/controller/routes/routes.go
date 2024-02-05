package routes

import (
  "github.com/gin-gonic/gin"
  "api_pdv/src/controller"
)

func InitRoutes(rg *gin.RouterGroup) {
  rg.POST("/product", controller.NewProduct)
  rg.GET("/product/:id", controller.GetProduct)
  rg.GET("/products", controller.GetProducts)
  rg.PUT("/product/:id", controller.UpdateProduct)
  rg.PUT("/sale/:id", controller.SaleProduct)
  rg.DELETE("/product/:id", controller.DeleteProduct)
}
