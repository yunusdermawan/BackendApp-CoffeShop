package handlers

import (
	"gogin/config"
	"gogin/internal/repositories"
	"gogin/pkg"

	"github.com/gin-gonic/gin"
)

type User struct {
	User_name     string `db:"user_name" json:"user_name" form:"user_name"`
	User_password string `db:"user_password" json:"user_password,omitempty"`
}

type HandlerAuth struct {
	*repositories.RepoUser
}

func NewAuth(r *repositories.RepoUser) *HandlerAuth {
	return &HandlerAuth{r}
}

func (h *HandlerAuth) Login(ctx *gin.Context) {
	var data User
	if ers := ctx.ShouldBind(&data); ers != nil {
		pkg.NewRes(500, &config.Result{
			Data: ers.Error(),
		}).Send(ctx)
		return
	}

	// if ers := ctx.ShouldBind(&data); ers != nil {
	// 	ctx.AbortWithError(http.StatusBadRequest, ers)
	// 	return
	// }

	users, err := h.GetAuthData(data.User_name)
	if err != nil {
		pkg.NewRes(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}
	// if err != nil {
	// 	ctx.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	if err := pkg.VerifyPassword(users.User_password, data.User_password); err != nil {
		pkg.NewRes(401, &config.Result{
			Data: "Password salah",
		}).Send(ctx)
		return
	}
	// if err := pkg.VerifyPassword(users.User_password, data.User_password); err != nil {
	// 	ctx.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	jwtt := pkg.NewToken(users.Id_user, users.Role)
	tokens, err := jwtt.Generate()
	if err != nil {
		pkg.NewRes(500, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}
	// if err != nil {
	// 	ctx.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	pkg.NewRes(200, &config.Result{Data: tokens}).Send(ctx)
	// ctx.JSON(200, tokens)
}
