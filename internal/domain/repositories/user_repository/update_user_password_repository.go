package user_repository

type UpdateUserPasswordRepository interface {
	UpdateUserPassword(userID int, newPassword string) error
}
