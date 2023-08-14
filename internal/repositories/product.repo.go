package repositories

import (
	"gogin/internal/models"
	"gogin/static"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type RepoProduct struct {
	*sqlx.DB
}

func NewProduct(db *sqlx.DB) *RepoProduct {
	return &RepoProduct{db}
}

func (r *RepoProduct) CreateProduct(data *models.Product) (string, error) {
	q := `
		INSERT INTO public.product(
			product_name,
			slug_product,
			product_banner,
			product_price,
			product_size,
			product_description,
			is_favorite,
			product_type
		)
		VALUES (
			:product_name,
			:slug_product,
			:product_banner,
			:product_price,
			:product_size,
			:product_description,
			:is_favorite,
			:product_type);
	`

	_, err := r.NamedExec(q, data)
	if err != nil {
		return "", err
	}

	return "1 data product created", nil
}

func (r *RepoProduct) DeleteProduct(data *models.Product) (map[string]interface{}, error) {
	q := `DELETE FROM public.product WHERE slug_product = :slug_product;`

	_, err := r.NamedExec(q, data)
	if err != nil {
		return nil, err
	}

	customStat := static.Response(200, "1 product deleted")

	return customStat, nil
}

func (r *RepoProduct) GetProduct() ([]models.Product, error) {
	// products := models.Product{}
	var products []models.Product
	q := `
		SELECT *
		FROM public.product
		ORDER BY created_at DESC;
	`

	err := r.DB.Select(&products, q)
	if err != nil {
		return nil, err
	}

	// return &config.Result{Data: products}, nil
	return products, nil
}

func (r *RepoProduct) SearchProduct(search *models.Search) (map[string]interface{}, error) {
	var products []models.Product

	q := `
		SELECT *
		FROM public.product
		WHERE product_name LIKE $1
		AND product_type = $2;
	`

	src := "%" + search.Prod_name + "%"
	typ := search.SortBy_Typ
	err := r.DB.Select(&products, q, src, typ)
	if err != nil {
		return nil, err
	}

	customStat := static.Response(200, "products")
	return customStat, nil
}

func (r *RepoProduct) GetProductByPage(page *models.Page) ([]models.Product, error) {
	var products []models.Product

	pageInt, err := strconv.Atoi(page.Page)
	if err != nil || pageInt <= 0 {
		pageInt = 1
	}

	pageSizeInt, err := strconv.Atoi(page.Limit)
	if err != nil || pageSizeInt <= 0 {
		pageSizeInt = 5
	}

	offset := (pageInt - 1) * pageSizeInt

	q := `
		SELECT *
		FROM public.product
		LIMIT $1 OFFSET $2;
	`

	err = r.DB.Select(&products, q, pageSizeInt, offset)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *RepoProduct) UpdateProduct(data *models.Product) (string, error) {
	q := `
		UPDATE public.product
		SET
			product_name=:product_name,
			slug_product=:slug_product,
			product_banner=:product_banner,
			product_price=:product_price,
			product_size=:product_size,
			product_description=:product_description,
			updated_at=:updated_at,
			is_favorite=:is_favorite,
			product_type=:product_type
		WHERE slug_product = :slug_product;
	`

	_, err := r.NamedExec(q, data)
	if err != nil {
		return "", err
	}

	return "1 product updated", nil
}
