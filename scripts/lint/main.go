package main

import (
	"os"

	"github.com/Phillezi/common/scripts/glint/cmd/cli"
)

func main() {
	if err := cli.ExecuteE(); err != nil {
		os.Exit(1)
	}
}
