package main

import (
	"book/config"
	"book/route"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	app := config.NewApp()

	router := gin.New()

	route.RegisterRoutes(route.Book{}, router)

	log, _ := strconv.ParseBool(app.Debug)

	if log {
		router.Use(gin.Logger())
	}

	fmt.Println("Starting server at: ", app.Url+":", app.Port)

	router.Run(":" + app.Port)

}
