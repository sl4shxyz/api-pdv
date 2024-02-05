package controller

import (
  "api_pdv/src/configuration/rest_error"
  "api_pdv/src/configuration/validation"
  "api_pdv/src/repository"
  "api_pdv/src/controller/model/request"
  "github.com/gin-gonic/gin"

  "net/http"
)

func SaleProduct(c *gin.Context) {
  id := c.Param("id")

  var saleRequest request.SaleRequest

  if err := c.ShouldBindJSON(&saleRequest); err != nil {
    msgErr := validation.ValidateUserError(err)
    c.JSON(msgErr.Code, msgErr)
    return
  }

  if _, err := repository.GetProduct(id); err != nil {
    msgErr := rest_error.NewNotFoundError("product not found")
    c.JSON(msgErr.Code, msgErr)
    return
  }

  if err := repository.SaleProduct(id, &saleRequest); err != nil {
    msgErr := rest_error.NewBadRequestError("error when updating sale")
    c.JSON(msgErr.Code, msgErr)
    return
  }

  c.JSON(http.StatusOK, gin.H{"message": "updated sale"})
}
