CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE public.user (
    id_user uuid NULL DEFAULT uuid_generate_v4(),
    user_name varchar(255) NULL,
	first_name varchar(255) NOT NULL,
	last_name varchar(255) NOT NULL,
	user_password varchar(255) NOT NULL,
	user_gender varchar(255) NOT NULL,
    user_slug varchar(255) NOT NULL UNIQUE,
    user_banner varchar(255) NULL,
    user_email varchar(255) NOT NULL,
	user_phone varchar(255) NOT NULL,
	user_address varchar(255) NOT NULL,
	user_dob date NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NULL,
    CONSTRAINT user_pk PRIMARY KEY (id_user)
);