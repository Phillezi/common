package lint

import (
	"os"

	"github.com/Phillezi/common/scripts/lint/cmd/cli"
)

func main() {
	if err := cli.ExecuteE(); err != nil {
		os.Exit(1)
	}
}
