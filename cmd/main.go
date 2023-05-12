package main

import (
	"fmt"
	"os"
	tasklist "tasklist-cli"
)

func main() {
	if err := tasklist.AddTask(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
