//go:build linux || darwin

package main

import (
	"fmt"

	sh "github.com/bitfield/script"
)

func Shutdown(t int64) {
	sh.Exec(fmt.Sprintf("shutdown -P %d", t)).Stdout()
}
