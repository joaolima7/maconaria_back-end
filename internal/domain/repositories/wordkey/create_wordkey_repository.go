package wordkey

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type CreateWordKeyRepository interface {
	CreateWordKey(wordkey *entity.WordKey) (*entity.WordKey, error)
}
