package repositories

import (
	"errors"
	"gogin/config"
	"gogin/internal/models"

	"github.com/jmoiron/sqlx"
)

type RepoUserIF interface {
	CreateUser(data *models.User) (*config.Result, error)
	DeleteUser(data *models.User) (*config.Result, error)
	GetUser() (*config.Result, error)
	SearchUser(search *models.Search) (*config.Result, error)
	UpdateUser(data *models.User) (*config.Result, error)
	GetAuthData(user string) (*models.User, error)
}
type RepoUser struct {
	*sqlx.DB
}

func NewUser(db *sqlx.DB) *RepoUser {
	return &RepoUser{db}
}

func (r *RepoUser) CreateUser(data *models.User) (*config.Result, error) {
	q := `
		INSERT INTO public.user(
			user_name,
			user_password,
			user_email,
			role
		)
		VALUES (
			:user_name,
			:user_password,
			:user_email,
			:role);
	`

	_, err := r.NamedExec(q, data)
	if err != nil {
		return nil, err
	}

	// return "1 data user added", nil
	// return data.Role, nil
	return &config.Result{Message: "1 data user added"}, nil
}

func (r *RepoUser) DeleteUser(data *models.User) (*config.Result, error) {
	q := `DELETE FROM public.user WHERE id_user = :id_user;`

	_, err := r.NamedExec(q, data)
	if err != nil {
		return nil, err
	}

	// return "1 user deleted", nil
	return &config.Result{Message: "1 user deleted"}, nil
}

func (r *RepoUser) GetUser() (*config.Result, error) {
	// users := models.User{}
	var users []models.User
	q := `
		SELECT *
		FROM public.user
		ORDER BY created_at DESC;
	`

	err := r.DB.Select(&users, q)
	if err != nil {
		return nil, err
	}

	return &config.Result{Data: users}, nil
	// return users, nil
}

func (r *RepoUser) SearchUser(search *models.Search) (*config.Result, error) {
	var users []models.User

	q := `
		SELECT *
		FROM public.user
		WHERE user_name LIKE $1
		AND first_name = $2;
	`

	src := "%" + search.Prod_name + "%"
	typ := search.SortBy_Typ
	err := r.DB.Select(&users, q, src, typ)
	if err != nil {
		return nil, err
	}

	// return users, nil
	return &config.Result{Data: users}, nil
}

func (r *RepoUser) UpdateUser(data *models.User) (*config.Result, error) {
	q := `
		UPDATE public.user
		SET
			user_name = :user_name,
			first_name = :first_name,
			last_name = :last_name,
			user_password = :user_password,
			user_gender = :user_gender,
			user_banner = :user_banner,
			user_email = :user_email,
			user_phone = :user_phone,
			user_address = :user_address,
			user_dob = :user_dob,
			role = :role
		WHERE id_user = :id_user;
	`

	_, err := r.NamedExec(q, data)
	if err != nil {
		return nil, err
	}

	// return "1 user updated", nil
	return &config.Result{Message: "1 user updated"}, nil
}

func (r *RepoUser) GetAuthData(user string) (*models.User, error) {
	var result models.User
	q := `SELECT id_user, user_name, "role", user_password FROM public.user WHERE user_name = ?`

	if err := r.Get(&result, r.Rebind(q), user); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New("username not found")
		}

		return nil, err
	}

	return &result, nil
}
