package worker_pool

// Foremaner represents foremans interface
type Foremaner interface {
	StartWork(countOfWorkers int)
}
