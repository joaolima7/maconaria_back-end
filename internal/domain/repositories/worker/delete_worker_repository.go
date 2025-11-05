package worker

type DeleteWorkerRepository interface {
	DeleteWorker(id string) error
}
