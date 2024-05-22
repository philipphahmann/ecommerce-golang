package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Product was created successfully"})

	// pegar as infos do body

	// dar um bindJSON no request struct --> validar os campos se estao corretos

	// enriquece o struct de schema Product com o request e outras infos

	// conexao com o database

	// salva o struct no database

	// retorna mensagem de sucesso caso nao ocorra nenhum erro

}
