package repo

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/suryasaputra2016/essentask/model"
)

type UserRepo struct {
	db *sql.DB
}

func NewuserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (ur UserRepo) Create(user *model.User) error {
	query := `
		INSERT INTO users (name, email, password_hash)
		VALUES $1, $2, $3, $4, $5
		RETURNING id;`
	row := ur.db.QueryRow(query, user.Name, user.Email, user.PasswordHash)
	err := row.Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("creating user: %w", err)
	}
	log.Println("user created.")
	return nil
}

func (ur UserRepo) GetByEmail(email string) (*model.User, error) {
	query := `
		SELECT * FROM users
		WHERE email = $1;`
	row := ur.db.QueryRow(query, email)
	var user model.User
	err := row.Scan(&user)
	if err != nil {
		return nil, fmt.Errorf("getting user by email: %w", err)
	}
	return &user, nil
}
