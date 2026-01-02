package main

import (
	"my/config"
	"my/database"
	"my/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()
	database.Connect(config.Get("DB_URL"))

	r := gin.Default()
	routes.Setup(r)

	r.Run(":" + config.Get("PORT"))
}
