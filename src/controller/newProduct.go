package controller

import (
  "api_pdv/src/configuration/validation"
  "github.com/gin-gonic/gin"
  "api_pdv/src/controller/model/request"
  "api_pdv/src/repository"
  "api_pdv/src/configuration/rest_error"

  "net/http"
)

func NewProduct(c *gin.Context) {
  var productRequest request.ProductRequest

  if err := c.ShouldBindJSON(&productRequest); err != nil {
    msgErr := validation.ValidateUserError(err)
    c.JSON(msgErr.Code, msgErr)
    return
  }


  if err := repository.CreateProduct(&productRequest); err != nil {
    msgErr := rest_error.NewInternalServerError("Error when saving product to database")
    c.JSON(msgErr.Code, msgErr)
    return
  }

  c.JSON(http.StatusOK, gin.H{"message": "product created"})
}
