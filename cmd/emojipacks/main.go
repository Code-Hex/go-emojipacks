package main

import (
	"os"

	"github.com/Code-Hex/go-emojipacks"
)

func main() {
	os.Exit(emojipacks.Run(os.Args))
}
