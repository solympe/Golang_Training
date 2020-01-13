package main

import wp "github.com/solympe/Golang_Training/pkg/worker-pool"

const countOfWorkers = 5

func main() {
	chanOfMaster := make(chan string, countOfWorkers)

	master := wp.NewForemaner(chanOfMaster)
	master.StartWork(countOfWorkers)

	close(chanOfMaster)
}
