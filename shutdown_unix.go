//go:build linux || darwin

package main

import sh "github.com/bitfield/script"

func Shutdown() {
	sh.Exec("shutdown -P now").Stdout()
}
