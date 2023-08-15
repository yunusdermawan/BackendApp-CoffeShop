package routers

import (
	"gogin/internal/handlers"
	"gogin/internal/middlewares"
	"gogin/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func user(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/user")

	repo := repositories.NewUser(d)
	handler := handlers.NewUser(repo)

	route.POST("/", middlewares.AuthJwt("admin"), handler.CreateData)
	route.GET("/all", middlewares.AuthJwt("admin", "user"), handler.GetData)
	route.GET("/search", middlewares.AuthJwt("admin", "user"), handler.SearchData)
	route.PATCH("/:id_user", middlewares.AuthJwt("admin"), handler.UpdateData)
	route.DELETE("/:id_user", middlewares.AuthJwt("admin"), handler.DeleteData)

}
