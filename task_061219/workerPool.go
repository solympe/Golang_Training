package main

import (
	"fmt"
	"math/rand"
	"time"
)

//написать программу worker pool с использованием каналов, входные параметры - количество рабочих
//у нас есть рабочие, которые начинают работу параллельно (время работы - рандомное число от 1 до 3 секунд),
//затем они завершают работу и сдают свои результаты на проверку прорабу, который их принимает

type prorab struct {
	channel chan string
}

func (p *prorab) checkWork() {
	for {
		rand.Seed(86)

		// need to change
		time.Sleep(3 * time.Second)
		p.channel <- "I ve done"
	}
}

func main() {
	countOfWorkers := 3
	chanOfProrab := make(chan string,countOfWorkers)
	proRab := prorab{channel:chanOfProrab}

	go proRab.checkWork()
	for countOfWorkers != 0{
		i :=  <- chanOfProrab
		fmt.Println(i)
		countOfWorkers--
	}
	defer close(chanOfProrab)
}