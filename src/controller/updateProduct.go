package controller

import (
  "api_pdv/src/repository"
  "api_pdv/src/configuration/validation"
  "api_pdv/src/configuration/rest_error"
  "api_pdv/src/controller/model/request"
  "github.com/gin-gonic/gin"

  "net/http"
)

func UpdateProduct(c *gin.Context) {
  id := c.Param("id")

  var productRequest request.ProductRequest

  if err := c.ShouldBindJSON(&productRequest); err != nil {
    msgErr := validation.ValidateUserError(err)
    c.JSON(msgErr.Code, msgErr)
    return
  }

  if _, err := repository.GetProduct(id); err != nil {
    msgErr := rest_error.NewNotFoundError("product not found")
    c.JSON(msgErr.Code, msgErr)
    return
  }

  if err := repository.UpdateProduct(id, &productRequest); err != nil {
    msgErr := rest_error.NewInternalServerError("error when updating product")
    c.JSON(msgErr.Code, msgErr)
    return
  }

  c.JSON(http.StatusOK, gin.H{"message": "updated product"})
}
