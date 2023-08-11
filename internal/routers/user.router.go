package routers

import (
	"gogin/internal/handlers"
	"gogin/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func user(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/user")

	repo := repositories.NewUser(d)
	handler := handlers.NewUser(repo)

	route.POST("/", handler.CreateData)
	route.GET("/all", handler.GetData)
	route.GET("/", handler.GetByPage)
	route.GET("/search", handler.SearchData)
	route.PATCH("/:slug", handler.UpdateData)
	route.DELETE("/:slug", handler.DeleteData)

}