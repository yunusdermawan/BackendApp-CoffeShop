package handlers

import (
	"gogin/internal/models"
	"gogin/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerProduct struct {
	*repositories.RepoProduct
}

func NewProduct(r *repositories.RepoProduct) *HandlerProduct {
	return &HandlerProduct{r}
}

func (h *HandlerProduct) CreateData(ctx *gin.Context) {
	var product models.Product

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := h.CreateProduct(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, response)
}

func (h *HandlerProduct) DeleteData(ctx *gin.Context) {
	var product models.Product

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	product.Slg_prod = ctx.Param("slug")
	response, err := h.DeleteProduct(&product)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, response)
}

func (h *HandlerProduct) GetData(ctx *gin.Context) {
	var product models.Product

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := h.GetProduct(&product)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, response)
}

func (h *HandlerProduct) SearchData(ctx *gin.Context) {
	var product models.Product

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := h.SearchProduct(&product)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, response)
}

func (h *HandlerProduct) UpdateData(ctx *gin.Context) {
	var product models.Product

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	product.Slg_prod = ctx.Param("slug")
	response, err := h.UpdateProduct(&product)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, response)
}
