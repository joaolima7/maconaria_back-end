package user

type UpdateUserPasswordRepository interface {
	UpdateUserPassword(userID string, newPassword string) error
}
