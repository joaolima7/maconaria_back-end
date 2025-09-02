package entity

import "github.com/joaolima7/maconaria_back-end/internal/types"

type Post struct {
	ID          string
	Title       string
	Description string
	Date        string
	Image       string
	User        *User
	PostType    types.PostType
}

func NewPost(id string, title string, description string, date string, image string, user User, postType types.PostType) (*Post, error) {
	return &Post{
		ID:          id,
		Title:       title,
		Description: description,
		Date:        date,
		Image:       image,
		User:        &user,
		PostType:    postType,
	}, nil
}
