package entity

import (
	"regexp"
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"golang.org/x/crypto/bcrypt"
)

type UserDegree string

const (
	DegreeApprentice UserDegree = "apprentice"
	DegreeCompanion  UserDegree = "companion"
	DegreeMaster     UserDegree = "master"
)

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	CIM       string
	Degree    UserDegree
	IsActive  bool
	IsAdmin   bool
	IsRegular bool
	Posts     []*Post
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(id string, name string, email string, password string, cim string, degree UserDegree, isActive bool, isAdmin bool, isRegular bool) (*User, error) {
	user := &User{
		ID:        id,
		Name:      name,
		Email:     email,
		Password:  password,
		CIM:       cim,
		Degree:    degree,
		IsActive:  isActive,
		IsAdmin:   isAdmin,
		IsRegular: isRegular,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := user.ValidateName(); err != nil {
		return nil, err
	}
	if err := user.ValidateEmail(); err != nil {
		return nil, err
	}
	if err := user.ValidateCIM(); err != nil {
		return nil, err
	}
	if err := user.ValidateDegree(); err != nil {
		return nil, err
	}
	if err := user.ValidatePassword(password); err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, apperrors.NewInternalError("Falha ao criptografar senha!", err)
	}
	user.Password = string(hash)

	return user, nil
}

func (u *User) ValidatePassword(password string) error {
	if len(password) < 8 {
		return apperrors.NewValidationError("senha", "A senha deve conter no mínimo 8 caracteres!")
	}
	return nil
}

func (u *User) ValidateEmail() error {
	if len(u.Email) == 0 {
		return apperrors.NewValidationError("e-mail", "O e-mail não pode ser vazio!")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(u.Email) {
		return apperrors.NewValidationError("e-mail", "Formato de e-mail inválido!")
	}
	return nil
}

func (u *User) ValidateName() error {
	if len(u.Name) == 0 {
		return apperrors.NewValidationError("nome", "O nome não pode ser vazio!")
	}
	if len(u.Name) < 3 {
		return apperrors.NewValidationError("nome", "O nome deve ter no mínimo 3 caracteres!")
	}
	return nil
}

func (u *User) ValidateCIM() error {
	if len(u.CIM) == 0 {
		return apperrors.NewValidationError("CIM", "O CIM não pode ser vazio!")
	}
	return nil
}

func (u *User) ValidateDegree() error {
	validDegrees := []UserDegree{DegreeApprentice, DegreeCompanion, DegreeMaster}
	for _, valid := range validDegrees {
		if u.Degree == valid {
			return nil
		}
	}
	return apperrors.NewValidationError("grau", "O grau deve ser 'apprentice', 'companion' ou 'master'!")
}

func (u *User) HashPassword(password string) error {
	if err := u.ValidatePassword(password); err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return apperrors.NewInternalError("Falha ao criptografar senha!", err)
	}
	u.Password = string(hash)
	return nil
}
