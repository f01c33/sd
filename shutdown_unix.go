//go:build linux || darwin

package main

import (
	"fmt"
	"time"

	sh "github.com/bitfield/script"
)

func Shutdown(t time.Time) {
	sh.Exec(fmt.Sprintf("shutdown +%d", int64(time.Until(t).Abs().Minutes()))).Stdout()
}
