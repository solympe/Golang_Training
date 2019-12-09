package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type prorab struct {
	channel chan string
}

func (p *prorab) checkWork(i int, wg *sync.WaitGroup) {
		fmt.Println("I am worker " + strconv.Itoa(i) + " and i start my job")
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(3)
		time.Sleep(time.Duration(n) * time.Second)
		p.channel <- "I am worker " + strconv.Itoa(i) + " and i finished my job in " + strconv.Itoa(n) + " seconds"
		wg.Done()
}

func main() {
	var wg sync.WaitGroup
	countOfWorkers := 3
	chanOfMaster := make(chan string,countOfWorkers)

	master := prorab{channel:chanOfMaster}

	for i := 1; i <= countOfWorkers; i++ {
		wg.Add(1)
		go master.checkWork(i, &wg)
	}
	wg.Wait()

	for i := 1; i <= countOfWorkers; i++ {
		i := <- chanOfMaster
	}

	defer close(chanOfMaster)
}
