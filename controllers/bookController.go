package controllers

import (
	"book/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookController struct {
}

func (b BookController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, models.All(models.Book{}))
}
func (b BookController) Store(c *gin.Context) {
	c.MultipartForm()
	c.JSON(http.StatusOK, models.Create(models.Book{}, c.Request.PostForm))
}
func (b BookController) Show(c *gin.Context) {
	c.JSON(http.StatusOK, models.Find(models.Book{}, c.Param("id")))
}
func (b BookController) Update(c *gin.Context) {
	c.MultipartForm()
	c.JSON(http.StatusOK, models.Update(models.Book{}, c.Param("id"), c.Request.PostForm))
}
func (b BookController) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, models.Delete(models.Book{}, c.Param("id")))
}
