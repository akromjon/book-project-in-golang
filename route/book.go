package route

import (
	"book/controllers"

	"github.com/gin-gonic/gin"
)

type Book struct {
}

func (b Book) Routes(router *gin.Engine) {

	router.GET("/book", func(ctx *gin.Context) {
		controllers.Index(controllers.BookController{}, ctx)
	})

	router.POST("/book", func(ctx *gin.Context) {
		controllers.Store(controllers.BookController{}, ctx)
	})

	router.GET("/book/:id", func(ctx *gin.Context) {
		controllers.Show(controllers.BookController{}, ctx)
	})

	router.PUT("/book/:id", func(ctx *gin.Context) {
		controllers.Update(controllers.BookController{}, ctx)
	})

	router.DELETE("/book/:id", func(ctx *gin.Context) {
		controllers.Delete(controllers.BookController{}, ctx)
	})
}
