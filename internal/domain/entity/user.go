package entity

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
	IsActive bool
	IsAdmin  bool
	Posts    []*Post
}

func NewUser(id string, name string, email string, password string, isActive bool, isAdmin bool) (*User, error) {
	user := &User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
		IsActive: isActive,
		IsAdmin:  isAdmin,
	}

	err := user.validatePassword()
	if err != nil {
		return nil, err
	}

	err = user.validateEmail()
	if err != nil {
		return nil, err
	}

	err = user.validateName()
	if err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("falha na criptografia de senha")
	}

	user.Password = string(hash)

	return user, nil
}

func (u *User) validatePassword() error {
	if len(u.Password) < 8 {
		return errors.New("a senha deve conter ao menos 8 caracteres")
	}

	return nil
}

func (u *User) validateEmail() error {
	if len(u.Email) == 0 {
		return errors.New("o email não pode ser vazio")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(u.Email) {
		return errors.New("email inválido")
	}

	return nil
}

func (u *User) validateName() error {
	if len(u.Name) == 0 {
		return errors.New("o nome não pode ser vazio")
	}

	return nil
}
