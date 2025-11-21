package entity

import "time"

type PostImage struct {
	ID        string
	PostID    string
	ImageURL  string
	CreatedAt time.Time
}

func NewPostImage(id string, postID string, imageURL string) *PostImage {
	return &PostImage{
		ID:        id,
		PostID:    postID,
		ImageURL:  imageURL,
		CreatedAt: time.Now(),
	}
}
