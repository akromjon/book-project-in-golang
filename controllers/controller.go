package controllers

import "github.com/gin-gonic/gin"

type Controller interface {
	Index(context *gin.Context)
	Store(context *gin.Context)
	Show(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

func Index(c Controller, context *gin.Context) {
	c.Index(context)
}
func Store(c Controller, context *gin.Context) {
	c.Store(context)
}
func Show(c Controller, context *gin.Context) {
	c.Show(context)
}
func Update(c Controller, context *gin.Context) {
	c.Update(context)
}
func Delete(c Controller, context *gin.Context) {
	c.Delete(context)
}
