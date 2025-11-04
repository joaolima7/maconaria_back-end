package user_usecase

import (
	"github.com/joaolima7/maconaria_back-end/internal/domain/apperrors"
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/domain/repositories/user"
	"golang.org/x/crypto/bcrypt"
)

type UpdateUserPasswordInputDTO struct {
	ID              string `json:"id"`
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required,min=8"`
}

type UpdateUserPasswordUseCase struct {
	UpdateRepository      user.UpdateUserPasswordRepository
	GetUserByIdRepository user.GetUserByIdRepository
}

func NewUpdateUserPasswordUseCase(
	updateRepository user.UpdateUserPasswordRepository,
	getUserByIdRepository user.GetUserByIdRepository,
) *UpdateUserPasswordUseCase {
	return &UpdateUserPasswordUseCase{
		UpdateRepository:      updateRepository,
		GetUserByIdRepository: getUserByIdRepository,
	}
}

func (uc *UpdateUserPasswordUseCase) Execute(input UpdateUserPasswordInputDTO) error {
	existingUser, err := uc.GetUserByIdRepository.GetUserById(input.ID)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(input.CurrentPassword)); err != nil {
		return apperrors.NewUnauthorizedError("Senha atual não corresponde!")
	}

	tempUser := &entity.User{}
	if err := tempUser.ValidatePassword(input.NewPassword); err != nil {
		return err
	}

	if err := tempUser.HashPassword(input.NewPassword); err != nil {
		return err
	}

	return uc.UpdateRepository.UpdateUserPassword(input.ID, tempUser.Password)
}
