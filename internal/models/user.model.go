package models

import (
	"time"
)

type User struct {
	Id_user       string     `db:"id_user" form:"id_user" json:"id_user" valid:"-"`
	User_name     string     `db:"user_name" form:"user_name" json:"user_name" valid:"type(string)"`
	First_name    *string    `db:"first_name" form:"first_name" json:"first_name" valid:"-"`
	Last_name     *string    `db:"last_name" form:"last_name" json:"last_name" valid:"-"`
	User_password string     `db:"user_password" form:"user_password" json:"user_password" valid:"stringlength(6|10)"`
	User_gender   *string    `db:"user_gender" form:"user_gender" json:"user_gender" valid:"-"`
	User_banner   *string    `db:"user_banner" form:"user_banner" json:"user_banner" valid:"-"`
	User_email    *string    `db:"user_email" form:"user_email" json:"user_email" valid:"-"`
	User_phone    *string    `db:"user_phone" form:"user_phone" json:"user_phone" valid:"-"`
	User_address  *string    `db:"user_address" form:"user_address" json:"user_address" valid:"-"`
	User_dob      *string    `db:"user_dob" form:"user_dob" json:"user_dob" valid:"-"`
	Crt_at        *time.Time `db:"created_at" json:"created_at" valid:"-"`
	Upd_at        *time.Time `db:"updated_at" json:"updated_at" valid:"-"`
	Role          string     `db:"role" form:"role" json:"role" valid:"type(string)"`
}
