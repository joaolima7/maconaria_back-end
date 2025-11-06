package library

type DeleteLibraryRepository interface {
	DeleteLibrary(id string) error
}
