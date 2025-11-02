package entity

import (
	"time"

	"github.com/joaolima7/maconaria_back-end/internal/types"
)

type Post struct {
	ID                  string
	Title               string
	Category            string
	SmallDescription    string
	CompleteDescription string
	Date                *string
	Time                *string
	Location            *string
	IsFeatured          bool
	PostType            types.PostType
	UserID              string
	User                *User
	Images              []*PostImage
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func NewPost(id string, title string, category string, smallDesc string, completeDesc string, date *string, timeStr *string, location *string, isFeatured bool, postType types.PostType, userID string, user *User) (*Post, error) {
	return &Post{
		ID:                  id,
		Title:               title,
		Category:            category,
		SmallDescription:    smallDesc,
		CompleteDescription: completeDesc,
		Date:                date,
		Time:                timeStr,
		Location:            location,
		IsFeatured:          isFeatured,
		PostType:            postType,
		UserID:              userID,
		User:                user,
		Images:              nil,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}, nil
}
