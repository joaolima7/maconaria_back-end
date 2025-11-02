package entity

import "time"

type PostImage struct {
	ID        string
	PostID    string
	ImageData []byte
	CreatedAt time.Time
}

func NewPostImage(id string, postID string, data []byte) *PostImage {
	return &PostImage{
		ID:        id,
		PostID:    postID,
		ImageData: data,
		CreatedAt: time.Now(),
	}
}
