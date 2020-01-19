package main

import foreman "github.com/solympe/Golang_Training/pkg/other/worker-pool"

const countOfWorkers = 5

func main() {
	chanOfMaster := make(chan string, countOfWorkers)

	master := foreman.NewForeman(chanOfMaster)
	master.StartWork(countOfWorkers)

	close(chanOfMaster)
}
