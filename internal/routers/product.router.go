package routers

import (
	"gogin/internal/handlers"
	"gogin/internal/middlewares"
	"gogin/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func product(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/product")

	repo := repositories.NewProduct(d)
	handler := handlers.NewProduct(repo)

	route.POST("/", middlewares.AuthJwt("admin", "user"), middlewares.UploadFile, handler.CreateData)
	route.GET("/all", handler.GetData)
	route.GET("/", handler.GetByPage)
	route.GET("/search", handler.SearchData)
	route.PATCH("/:slug", middlewares.AuthJwt("admin", "user"), middlewares.UploadFile, handler.UpdateData)
	route.DELETE("/:slug", middlewares.AuthJwt("admin"), handler.DeleteData)

}
