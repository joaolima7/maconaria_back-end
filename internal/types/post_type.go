package types

type PostType string

const (
	PostTypeEvent         PostType = "event"
	PostTypeArticle       PostType = "article"
	PostTypeCommemoration PostType = "commemoration"
	PostTypeNews          PostType = "news"
)
