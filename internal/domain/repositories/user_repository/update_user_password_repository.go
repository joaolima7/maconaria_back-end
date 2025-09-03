package user_repository

type UpdateUserPasswordRepository interface {
	UpdateUserPassword(userID string, newPassword string) error
}
