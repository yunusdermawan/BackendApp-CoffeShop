package handlers

import (
	"gogin/internal/models"
	"gogin/internal/repositories"
	"gogin/pkg"
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
	product.Slg_prod = ctx.Param("slug")

	response, err := h.DeleteProduct(&product)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, response)
}

func (h *HandlerProduct) GetData(ctx *gin.Context) {
	data, err := h.GetProduct()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	pkg.NewRes(200, data).Send(ctx)
}

func (h *HandlerProduct) SearchData(ctx *gin.Context) {
	var search models.Search

	if err := ctx.ShouldBindQuery(&search); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := h.SearchProduct(&search)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, response)
}

func (h *HandlerProduct) GetByPage(ctx *gin.Context) {
	var page models.Page

	if err := ctx.ShouldBindQuery(&page); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := h.GetProductByPage(&page)

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
