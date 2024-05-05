package main

import sh "github.com/bitfield/script"

func Shutdown() {
	sh.Exec("shutdown /s /p now").Stdout()
}
