package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	k, err := parseArg(os.Args[1:])

	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	handleTicker(k)

	<-sigs
	fmt.Println("Termination")
	close(sigs)
}

func parseArg(args []string) (uint, error) {
	if len(args) == 0 {
		return 0, fmt.Errorf("no args were passed")
	} else if len(args) != 1 {
		return 0, fmt.Errorf("1 arguement is wanted, passed instead %q", args)
	}

	k, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil {
		return 0, fmt.Errorf("arguement should be of type unsigned integer, instead %q", args[0])
	}

	return uint(k), nil
}

func handleTicker(k uint) {
	go func() {
		counter := 1
		for {
			time.Sleep(time.Duration(k) * time.Second)

			fmt.Printf("Tick %v since %v\n", counter, counter*int(k))

			counter++
		}
	}()
}
