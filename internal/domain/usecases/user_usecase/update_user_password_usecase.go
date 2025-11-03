package user_usecase

import (
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/user"
)

type UpdateUserPasswordInputDTO struct {
	ID          string `json:"id"`
	NewPassword string `json:"new_password" validate:"required,min=8"`
}

type UpdateUserPasswordUseCase struct {
	Repository user.UpdateUserPasswordRepository
}

func NewUpdateUserPasswordUseCase(repository user.UpdateUserPasswordRepository) *UpdateUserPasswordUseCase {
	return &UpdateUserPasswordUseCase{
		Repository: repository,
	}
}

func (uc *UpdateUserPasswordUseCase) Execute(input UpdateUserPasswordInputDTO) error {

	tempUser := &entity.User{}
	if err := tempUser.ValidatePassword(input.NewPassword); err != nil {
		return err
	}

	if err := tempUser.HashPassword(input.NewPassword); err != nil {
		return err
	}

	return uc.Repository.UpdateUserPassword(input.ID, tempUser.Password)
}
