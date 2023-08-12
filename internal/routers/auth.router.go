package routers

import (
	"gogin/internal/handlers"
	"gogin/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func auth(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/auth")

	repo := repositories.NewUser(d)
	handler := handlers.NewAuth(repo)

	route.POST("/login", handler.Login)

}
