package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id, "message": "Product was updated successfully"})
}
