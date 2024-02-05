package controller

import (
  "github.com/gin-gonic/gin"
  "api_pdv/src/configuration/rest_error"
  "api_pdv/src/repository"

  "net/http"
)

func DeleteProduct(c *gin.Context) {
  id := c.Param("id")

  if _, err := repository.GetProduct(id); err != nil {
    msgErr := rest_error.NewNotFoundError("product not found")
    c.JSON(msgErr.Code, msgErr)
    return
  }

  if err := repository.DeleteProduct(id); err != nil {
    msgErr := rest_error.NewBadRequestError("error when deleting product")
    c.JSON(msgErr.Code, msgErr)
    return
  }

  c.JSON(http.StatusOK, gin.H{"message": "deleted product"})
}
