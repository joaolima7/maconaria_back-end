package wordkey

type DeleteWordKeyRepository interface {
	DeleteWordKey(id string) error
}
