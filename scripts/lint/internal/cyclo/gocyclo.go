package cyclo

import (
	"fmt"

	"github.com/Phillezi/common/scripts/glint/internal/api"
	"github.com/Phillezi/common/scripts/glint/internal/defaults"
	"github.com/Phillezi/common/utils/or"
	"github.com/fzipp/gocyclo"
)

type Command struct {
	Threshold int
}

func (c *Command) Run(path string, options ...api.CommandOption) error {
	// Apply options (like threshold override)
	opts := &api.Options{
		Logger: api.DefaultLogger,
	}
	for _, opt := range options {
		opt(opts)
	}

	threshold := or.Or(c.Threshold, defaults.DefaultGocycloOverValue)

	opts.Logger.Infof("Running gocyclo with threshold > %d on %s", threshold, path)

	stats := gocyclo.Analyze([]string{path}, nil)

	var found bool
	for _, stat := range stats {
		if stat.Complexity > threshold {
			if !found {
				opts.Logger.Errorf("gocyclo found functions exceeding complexity threshold:")
				found = true
			}
			opts.Logger.PrintIndentf("[%s.%s]\t(complexity %d):\t%s:%d:%d",
				stat.PkgName, stat.FuncName, stat.Complexity, stat.Pos.Filename, stat.Pos.Line, stat.Pos.Offset)
		}
	}

	if found {
		return fmt.Errorf("gocyclo: complexity threshold (%d) exceeded", threshold)
	}

	opts.Logger.Successf("passed gocyclo -over %d", threshold)
	return nil
}
