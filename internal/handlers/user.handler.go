package handlers

import (
	"gogin/internal/models"
	"gogin/internal/repositories"
	"gogin/pkg"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	*repositories.RepoUser
}

func NewUser(r *repositories.RepoUser) *HandlerUser {
	return &HandlerUser{r}
}

func (h *HandlerUser) CreateData(ctx *gin.Context) {
	var users models.User
	var ers error

	if err := ctx.ShouldBind(&users); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//Payload validation
	_, ers = govalidator.ValidateStruct(&users)
	if ers != nil {
		ctx.AbortWithError(http.StatusBadRequest, ers)
		return
	}

	//Hash password
	users.User_password, ers = pkg.HashPassword(users.User_password)
	if ers != nil {
		ctx.AbortWithError(http.StatusBadRequest, ers)
		return
	}

	response, err := h.CreateUser(&users)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, response)
}

func (h *HandlerUser) DeleteData(ctx *gin.Context) {
	var users models.User
	users.Id_user = ctx.Param("id_user")

	response, err := h.DeleteUser(&users)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, response)
}

func (h *HandlerUser) GetData(ctx *gin.Context) {
	data, err := h.GetUser()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// pkg.NewRes(200, data).Send(ctx)
	ctx.JSON(200, data)
}

func (h *HandlerUser) SearchData(ctx *gin.Context) {
	var search models.Search

	if err := ctx.ShouldBindQuery(&search); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := h.SearchUser(&search)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, response)
}

func (h *HandlerUser) UpdateData(ctx *gin.Context) {
	var users models.User

	if err := ctx.ShouldBind(&users); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	users.Id_user = ctx.Param("id_user")
	response, err := h.UpdateUser(&users)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, response)
}
