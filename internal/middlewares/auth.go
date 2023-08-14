package middlewares

import (
	"gogin/config"
	"gogin/pkg"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthJwt(role ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var valid bool
		var header string

		if header = ctx.GetHeader("Authorization"); header == "" {
			pkg.NewRes(401, &config.Result{
				Message: "Please login",
			}).Send(ctx)
			return
		}
		// if header = ctx.GetHeader("Authorization"); header == "" {
		// 	ctx.JSON(200, "Please login")
		// 	return
		// }

		if !strings.Contains(header, "Bearer") {
			pkg.NewRes(401, &config.Result{
				Message: "Invalid header value",
			}).Send(ctx)
			return
		}
		// if !strings.Contains(header, "Bearer") {
		// 	ctx.JSON(200, "Invalid header value")
		// 	return
		// }

		tokens := strings.Replace(header, "Bearer ", "", -1)
		check, err := pkg.VerifyToken(tokens)
		if err != nil {
			pkg.NewRes(401, &config.Result{
				Message: err.Error(),
			}).Send(ctx)
			return
		}
		// if err != nil {
		// 	ctx.JSON(200, err.Error())
		// 	return
		// }

		for _, r := range role {
			if r == check.Role {
				valid = true
			}
		}

		if !valid {
			pkg.NewRes(401, &config.Result{
				Data: "you don't have permission to access",
			}).Send(ctx)
			return
		}
		// if !valid {
		// 	ctx.JSON(200, "You don't have permission to access")
		// 	return
		// }

		ctx.Set("userId", check.Id)
		ctx.Next()

	}
}
