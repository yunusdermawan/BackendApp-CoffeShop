package routers

import (
	"gogin/internal/handlers"
	"gogin/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func product(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/product")

	repo := repositories.NewProduct(d)
	handler := handlers.NewProduct(repo)

	route.POST("/", handler.CreateData)
	// route.GET("/", handler.GetData)
	route.GET("/", handler.GetByPage)
	route.GET("/search", handler.SearchData)
	route.PATCH("/:slug", handler.UpdateData)
	route.DELETE("/:slug", handler.DeleteData)

}
