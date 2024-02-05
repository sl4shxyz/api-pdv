package controller

import (
  "api_pdv/src/controller/model/response"
  "api_pdv/src/repository"
  "api_pdv/src/configuration/rest_error"
  "github.com/gin-gonic/gin"

  "net/http"
)

func GetProducts(c *gin.Context) {
  products, err := repository.GetProducts()
  if err != nil {
    msgErr := rest_error.NewBadRequestError("failed to get products")
    c.JSON(msgErr.Code, msgErr)
    return
  }

  var productResponse []response.ProductResponse

  for _, product := range products {
    productResponses := response.ProductResponse {
      ID: product.ID,
      Name: product.Name,
      Price: product.Price,
      Amount: product.Amount,
    }

    productResponse = append(productResponse, productResponses)
  }

  c.JSON(http.StatusOK, productResponse)
}

func GetProduct(c *gin.Context) {
  id := c.Param("id")

  product, err := repository.GetProduct(id)
  if err != nil {
    msgErr := rest_error.NewNotFoundError("product not found")
    c.JSON(msgErr.Code, msgErr)
    return
  }

  c.JSON(http.StatusOK, product)
}
