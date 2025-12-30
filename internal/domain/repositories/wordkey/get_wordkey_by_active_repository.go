package wordkey

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type GetWordKeyByActiveRepository interface {
	GetWordKeyByActive() (*entity.WordKey, error)
}
