package wordkey

import "github.com/joaolima7/maconaria_back-end/internal/domain/entity"

type GetAllWordKeysRepository interface {
	GetAllWordKeys() ([]*entity.WordKey, error)
}
