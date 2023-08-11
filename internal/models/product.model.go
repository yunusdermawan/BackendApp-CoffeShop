package models

import "time"

type Product struct {
	Id_prod   string     `db:"id_product" form:"id_product" json:"id_product"`
	Prod_name string     `db:"product_name" form:"product_name" json:"product_name"`
	Slg_prod  string     `db:"slug_product" form:"slug_product" json:"slug_product"`
	Prod_bnr  string     `db:"product_banner" form:"product_banner" json:"product_banner"`
	Prod_prc  string     `db:"product_price" form:"product_price" json:"product_price"`
	Prod_sz   string     `db:"product_size" form:"product_size" json:"product_size"`
	Prod_desc string     `db:"product_description" form:"product_description" json:"product_description"`
	Crt_at    *time.Time `db:"created_at" json:"created_at"`
	Upd_at    *time.Time `db:"updated_at" json:"updated_at"`
	Is_fav    bool       `db:"is_favorite" form:"is_favorite" json:"is_favorite"`
	Prod_tp   string     `db:"product_type" form:"product_type" json:"product_type"`
}

type Search struct {
	Prod_name  string `form:"product_name" uri:"product_name"`
	SortBy_Typ string `form:"sortBy_productType" uri:"sortBy_productType"`
}

type Page struct {
	Page  string `form:"page" uri:"page"`
	Limit string `form:"limit" uri:"limit"`
}
