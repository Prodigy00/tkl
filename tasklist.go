package tasklist

import (
	"io"
	"log"
)

type Runner interface {
	Init([]string) error
	Run() error
	Name() string
}

func Welcome(w io.Writer) {
	welcomeMsg := "Welcome to tasklist!"
	_, err := w.Write([]byte(welcomeMsg))
	if err != nil {
		log.Fatal(err)
	}
}
