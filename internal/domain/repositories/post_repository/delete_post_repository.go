package post_repository

type DeletePostRepository interface {
	Delete(postID string) error
}
