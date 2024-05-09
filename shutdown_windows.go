package main

import (
	"fmt"

	sh "github.com/bitfield/script"
)

func Shutdown(t int64) {
	sh.Exec(fmt.Sprintf("shutdown /s /p %d", t)).Stdout()
}
