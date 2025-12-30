package wordkey

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type GetWordKeyByIDRepository interface {
	GetWordKeyByID(id string) (*entity.WordKey, error)
}
