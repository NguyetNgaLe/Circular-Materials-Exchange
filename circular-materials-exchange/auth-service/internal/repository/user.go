package repository

import (
	"database/sql"
	"time"
)

type User struct {
	ID           string
	Name         string
	Email        string
	Phone        string
	PasswordHash string
	Role         string
	Avatar       string
	CompanyID    *string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *User) error {
	_, err := r.db.Exec(
		`INSERT INTO users (id, name, email, phone, password_hash, role, avatar, company_id)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		user.ID, user.Name, user.Email, user.Phone, user.PasswordHash, user.Role, user.Avatar, user.CompanyID,
	)
	return err
}

func (r *UserRepository) FindByEmail(email string) (*User, error) {
	user := &User{}
	err := r.db.QueryRow(
		`SELECT id, name, email, phone, password_hash, role, avatar, company_id, created_at, updated_at
		 FROM users WHERE email = $1`, email,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.PasswordHash, &user.Role, &user.Avatar, &user.CompanyID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindByID(id string) (*User, error) {
	user := &User{}
	err := r.db.QueryRow(
		`SELECT id, name, email, phone, password_hash, role, avatar, company_id, created_at, updated_at
		 FROM users WHERE id = $1`, id,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.PasswordHash, &user.Role, &user.Avatar, &user.CompanyID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Update(user *User) error {
	_, err := r.db.Exec(
		`UPDATE users SET name=$1, phone=$2, avatar=$3, updated_at=$4 WHERE id=$5`,
		user.Name, user.Phone, user.Avatar, time.Now(), user.ID,
	)
	return err
}
