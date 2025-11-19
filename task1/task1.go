package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"sort"
	"strconv"
	"sync"
	"time"
)

type goroutine struct {
	idx      int
	duration int
}

func main() {
	args, err := parseArguements()

	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	goroutineList := makeGoroutines(args[0], args[1])

	for i := range goroutineList {
		fmt.Printf("<%v, %v>\n", goroutineList[i].idx, goroutineList[i].duration)
	}
}

func parseArguements() ([]int, error) {
	args := os.Args[1:]

	if len(args) == 0 {
		return []int{}, fmt.Errorf("no args were passed")
	} else if len(args) != 2 {
		return []int{}, fmt.Errorf("2 arguements are wanted, passed instead %q", args)
	}

	n, err := strconv.Atoi(args[0])
	if err != nil {
		return []int{}, fmt.Errorf("first arguement should be of type integer, instead %q", args[0])
	} else if n < 0 {
		return []int{}, fmt.Errorf("first arguement should be positive, instead %q", args[0])
	}

	m, err := strconv.Atoi(args[1])
	if err != nil {
		return []int{}, fmt.Errorf("second arguement should be of type integer, instead %q", args[1])
	} else if m < 0 {
		return []int{}, fmt.Errorf("second arguement should be positive, instead %q", args[1])
	}

	return []int{n, m}, nil
}

func makeGoroutines(n, m int) []goroutine {
	goroutineList := make([]goroutine, n)

	var wg sync.WaitGroup
	for i := range goroutineList {
		duration := getRandomNum(m)

		wg.Go(func() {
			time.Sleep(time.Duration(duration) * time.Millisecond)
			goroutineList[i] = goroutine{i, duration}
		})
	}

	wg.Wait()

	sortGoroutines(goroutineList)

	return goroutineList
}

func getRandomNum(num int) int {
	return rand.IntN(num + 1)
}

func sortGoroutines(list []goroutine) {
	sort.Slice(list, func(i, j int) bool {
		return list[i].duration > list[j].duration
	})
}
