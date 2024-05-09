package main

import (
	"os"
	"strings"
	"time"

	nd "github.com/tj/go-naturaldate"
)

func main() {
	args := strings.Join(os.Args[1:], " ")
	t, err := nd.Parse(args, time.Now())
	if err != nil {
		panic(err)
	}
	Shutdown(t)
}
