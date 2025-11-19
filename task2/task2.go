package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args, err := parseArguements()

	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	channel := square(generator(args[0], args[1]))

	for i := range channel {
		fmt.Println(i)
	}
}

func parseArguements() ([]int, error) {
	args := os.Args[1:]

	if len(args) == 0 {
		return []int{}, fmt.Errorf("no args were passed")
	} else if len(args) != 2 {
		return []int{}, fmt.Errorf("2 arguements are wanted, passed instead %q", args)
	}

	k, err := strconv.Atoi(args[0])
	if err != nil {
		return []int{}, fmt.Errorf("first arguement should be of type integer, instead %q", args[0])
	}

	n, err := strconv.Atoi(args[1])
	if err != nil {
		return []int{}, fmt.Errorf("second arguement should be of type integer, instead %q", args[1])
	}

	return []int{k, n}, nil
}

func generator(k, n int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		if k <= n {
			for i := k; i <= n; i++ {
				out <- i
			}
		} else {
			for i := k; i >= n; i-- {
				out <- i
			}
		}
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := range in {
			out <- i * i
		}
	}()

	return out
}
