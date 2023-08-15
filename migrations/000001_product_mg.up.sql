CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE public.product (
    id_product uuid NULL DEFAULT uuid_generate_v4(),
    product_name varchar(255) NOT NULL,
    slug_product varchar(255) NOT NULL UNIQUE,
    product_banner varchar(255) NOT NULL,
    product_price varchar(255) NOT NULL,
	product_size varchar(255) NOT NULL,
	product_description varchar(500) NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NULL,
    CONSTRAINT product_pk PRIMARY KEY (id_product)
);