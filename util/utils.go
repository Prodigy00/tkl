package util

import (
	"bufio"
	"log"
	"os"
)

func CloseFile(f *os.File) {
	closeErr := f.Close()
	if closeErr != nil {
		log.Fatalf("error occurred closing the file. Err: %v", closeErr)
	}
}

func ReadTasksFromFile(f *os.File) (ts []string, err error) {
	var taskslice []string
	reader := bufio.NewScanner(f)

	for reader.Scan() {
		taskslice = append(taskslice, reader.Text())
	}

	if scanErr := reader.Err(); scanErr != nil {
		return nil, scanErr
	}
	return taskslice, nil
}
