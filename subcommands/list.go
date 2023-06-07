package subcommands

import (
	"flag"
	"fmt"
	"github.com/enescakir/emoji"
	"io"
	"log"
	"os"
)

const buf = 1024

func List(args []string) {
	flag := flag.NewFlagSet("list", flag.ExitOnError)

	err := flag.Parse(args)
	if err != nil {
		log.Fatal(err)
	}
	//list all tasks in chunks
	readFile("tklfile.txt")
}

func readFile(filename string) {
	//if file doesn't exist create it
	//if file is empty, print help message for adding tasks instead
	//else list tasks
	f, err := os.OpenFile("tklfile.txt", os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("error opening task file.")
	}

	fi, err := f.Stat()
	if err != nil {
		log.Fatal("error getting file size")
	}

	if fi.Size() == 0 {
		log.Printf("no tasks added, try adding some tasks %s", emoji.Rocket)
		NewHelp().PrintSubcommandHelp("add")
		os.Exit(1)
	}

	buffer := make([]byte, buf)

	for {
		bytesRead, err := f.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Failed to read file: %s", err)
		}

		if bytesRead == 0 {
			break
		}

		fmt.Print(string(buffer[:bytesRead]))
	}
}
