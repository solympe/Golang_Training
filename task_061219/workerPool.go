package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//написать программу worker pool с использованием каналов, входные параметры - количество рабочих
//у нас есть рабочие, которые начинают работу параллельно (время работы - рандомное число от 1 до 3 секунд),
//затем они завершают работу и сдают свои результаты на проверку прорабу, который их принимает

type prorab struct {
	channel chan string
}

func (p *prorab) checkWork(i int) {
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(3)
		time.Sleep(time.Duration(n) * time.Second)
		p.channel <- "I am worker " + strconv.Itoa(i) + " and i finished my job in " + strconv.Itoa(n) + " seconds"

	}
}

func main() {
	countOfWorkers := 3
	chanOfProrab := make(chan string,countOfWorkers)

	proRab := prorab{channel:chanOfProrab}

	for i := 1; i <= countOfWorkers; i++ {
	go proRab.checkWork(i)
	}

	for i := 1; i <= countOfWorkers; i++ {
		i := <- chanOfProrab
		fmt.Println(i)
	}

	defer close(chanOfProrab)
}