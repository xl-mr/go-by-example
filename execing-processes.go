package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	binary, err := exec.LookPath("ls")
	if err != nil {
		panic(err)
	}

	args :=[]string{"ls", "-alh"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
