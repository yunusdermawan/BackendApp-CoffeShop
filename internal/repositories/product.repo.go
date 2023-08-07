package repositories

import (
	"gogin/internal/models"

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

func (r *RepoProduct) DeleteProduct(data *models.Product) (string, error) {
	q := `DELETE FROM public.product WHERE slug_product = :slug_product;`

	_, err := r.NamedExec(q, data)
	if err != nil {
		return "", err
	}

	return "1 product deleted", nil
}

func (r *RepoProduct) GetProduct(data *models.Product) ([]models.Product, error) {
	var products []models.Product

	q := `
		SELECT 
			id_product,
			product_name,
			product_banner,
			product_price,
			product_size,
			product_description,
			is_favorite,
			product_type
		FROM public.product;
	`

	err := r.DB.Select(&products, q)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *RepoProduct) SearchProduct(search *models.Search) ([]models.Product, error) {
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

	return products, nil
}

func (r *RepoProduct) UpdateProduct(data *models.Product) (string, error) {
	q := `
		UPDATE public.product
		SET
			product_name=:product_name,
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
