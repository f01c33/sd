package main

import (
	"fmt"
	"time"

	sh "github.com/bitfield/script"
)

func Shutdown(t time.Time) {
	_, err := sh.Exec(fmt.Sprintf("shutdown /s /f %d", int64(time.Until(t).Abs().Minutes()))).Stdout()
	if err != nil {
		panic(err)
	}
}
