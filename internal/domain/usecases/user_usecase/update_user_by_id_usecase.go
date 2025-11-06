package user_usecase

import (
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/user"
)

type UpdateUserByIdInputDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	CIM      string `json:"cim" validate:"required"`
	Degree   string `json:"degree" validate:"required,oneof=apprentice companion master"`
	IsActive bool   `json:"is_active"`
	IsAdmin  bool   `json:"is_admin"`
}

type UpdateUserByIdOutputDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	CIM      string `json:"cim"`
	Degree   string `json:"degree"`
	IsActive bool   `json:"is_active"`
	IsAdmin  bool   `json:"is_admin"`
}

type UpdateUserByIdUseCase struct {
	Repository user.UpdateUserByIDRepository
}

func NewUpdateUserByIdUseCase(repository user.UpdateUserByIDRepository) *UpdateUserByIdUseCase {
	return &UpdateUserByIdUseCase{
		Repository: repository,
	}
}

func (uc *UpdateUserByIdUseCase) Execute(input UpdateUserByIdInputDTO) (*UpdateUserByIdOutputDTO, error) {
	user := &entity.User{
		ID:        input.ID,
		Name:      input.Name,
		Email:     input.Email,
		CIM:       input.CIM,
		Degree:    entity.UserDegree(input.Degree),
		IsActive:  input.IsActive,
		IsAdmin:   input.IsAdmin,
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

	updatedUser, err := uc.Repository.UpdateUserByID(user)
	if err != nil {
		return nil, err
	}

	return &UpdateUserByIdOutputDTO{
		ID:       updatedUser.ID,
		Name:     updatedUser.Name,
		Email:    updatedUser.Email,
		CIM:      updatedUser.CIM,
		Degree:   string(updatedUser.Degree),
		IsActive: updatedUser.IsActive,
		IsAdmin:  updatedUser.IsAdmin,
	}, nil
}
