package config

import (
	"github.com/gin-gonic/gin"
	"product-api/app/controllers/root"
	"product-api/config/collection"
	"product-api/db"
)

var Routers = gin.Default()

func init() {
	corsConfig(Routers)
	Routers.Static("/assets", "./assets")

	Routers.GET("/", root.Index)
	main := Routers.Group("v1")
	collection.MainRouter(db.DB, main)
}
