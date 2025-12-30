package mappers

import (
	"github.com/joaolima7/maconaria_back-end/internal/domain/entity"
	"github.com/joaolima7/maconaria_back-end/internal/infra/database/db"
)

func DbWordKeyToEntity(wordkeyDB db.Wordkey) *entity.WordKey {
	return &entity.WordKey{
		ID:        wordkeyDB.ID,
		WordKey:   wordkeyDB.Wordkey,
		Active:    wordkeyDB.Active,
		CreatedAt: wordkeyDB.CreatedAt.Time,
	}
}
