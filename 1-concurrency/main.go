package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

func main() {
	mainChan := make(chan float64)
	intChan := make(chan int)

	go makeRandomInts(intChan)
	go getPowInt(intChan, mainChan)
	for result := range mainChan {
		fmt.Printf("%d ", int(result))
	}
}

func makeRandomInts(intChan chan int) {
	numbers := make([]int, 10)

	for i := 0; i < len(numbers); i++ {
		numbers[i] = rand.IntN(101)
	}

	for _, number := range numbers {
		intChan <- number
	}
	close(intChan)
}

func getPowInt(intChan chan int, mainChan chan float64) {
	for randInt := range intChan {
		mainChan <- math.Pow(float64(randInt), 2)
	}
	close(mainChan)
}
