package wordkey

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type UpdateWordKeyByIDRepository interface {
	UpdateWordKeyByID(wordkey *entity.WordKey) (*entity.WordKey, error)
}
