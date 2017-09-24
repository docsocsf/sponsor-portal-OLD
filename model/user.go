package model

import (
	"database/sql"
	"errors"
	"log"

	"github.com/docsocsf/sponsor-portal/auth"
)

type User struct {
	Id   int64     `json:"id"`
	Name string    `json:"name"`
	Auth *UserAuth `json:"auth,omitempty"`
}

type UserAuth struct {
	Email string `json:"email,omitempty"`
}

type UserReader interface {
	Get(User) (User, error)
	GetById(auth.UserIdentifier) (User, error)
	HashedPassword(email string) (string, error)
}

const (
	getUserByEmail = `
	SELECT id FROM users WHERE email=$1
	`

	getUserById = `
	SELECT id, name, email FROM users WHERE id=$1
	`

	getHashedPassword = `
	SELECT hashed_password FROM users WHERE email=$1
	`
)

type userImpl struct {
	db *sql.DB
}

func NewUserReader(db *sql.DB) UserReader {
	return userImpl{db}
}

func (u userImpl) GetById(id auth.UserIdentifier) (User, error) {
	user := User{Auth: &UserAuth{}}
	err := u.db.QueryRow(getUserById, id.User).Scan(&user.Id, &user.Name, &user.Auth.Email)

	switch {
	case err == sql.ErrNoRows:
		return User{}, errors.New("User does not exist")
	case err != nil:
		return User{}, err
	default:
		return user, nil
	}
}

func (u userImpl) Get(user User) (User, error) {
	err := u.db.QueryRow(getUserByEmail, user.Auth.Email).Scan(&user.Id)

	switch {
	case err == sql.ErrNoRows:
		return User{}, errors.New("User does not exist")
	case err != nil:
		return User{}, err
	default:
		return user, nil
	}
}

func (u userImpl) HashedPassword(email string) (string, error) {
	log.Printf("Checking password for '%s'\n", email)

	var password string
	err := u.db.QueryRow(getHashedPassword, email).Scan(&password)

	switch {
	case err == sql.ErrNoRows:
		return "", errors.New("User does not exist")
	case err != nil:
		return "", err
	default:
		return password, nil
	}
}
