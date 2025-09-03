package user_usecase

import (
	"github.com/google/uuid"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/user_repository"
)

type UpdateUserByIdInputDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
	IsAdmin  bool   `json:"is_admin"`
}

type UpdateUserByIdOutputDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
	IsAdmin  bool   `json:"is_admin"`
}

type UpdateUserByIdUseCase struct {
	Repository user_repository.UpdateUserByIDRepository
}

func NewUpdateUserByIdUseCase(repository user_repository.UpdateUserByIDRepository) *UpdateUserByIdUseCase {
	return &UpdateUserByIdUseCase{
		Repository: repository,
	}
}

func (uc *UpdateUserByIdUseCase) Execute(input UpdateUserByIdInputDTO) (*UpdateUserByIdOutputDTO, error) {
	user, err := entity.NewUser(
		uuid.NewString(),
		input.Name,
		input.Email,
		"",
		input.IsActive,
		input.IsAdmin,
	)
	if err != nil {
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
		IsActive: updatedUser.IsActive,
		IsAdmin:  updatedUser.IsAdmin,
	}, nil
}
