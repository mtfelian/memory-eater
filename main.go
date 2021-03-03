package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("specify amount pls, in GB")
		fmt.Println("example: ", os.Args[0], "5")
		return
	}

	amount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	fmt.Printf("I will eat %d GB of your memory!\n", amount)

	b := make([]byte, amount<<30)
	for i := range b {
		b[i] = 0
	}

	fmt.Println("Waiting for Ctrl-C...")
	time.Sleep(7 * 24 * time.Hour)
	fmt.Println(len(b))
}
