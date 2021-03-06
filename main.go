package main

import (
	"fmt"
	"os"

	"github.com/kutifyhq/autifyctl/cmd"
)

const (
	exitOK = iota
	exitError
)

func main() {
	os.Exit(realmain(os.Args))
}

func realmain(args []string) int {
	if err := cmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		return exitError
	}

	return exitOK
}
