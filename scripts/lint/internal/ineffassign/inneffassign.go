package ineffassign

import (
	"fmt"

	"github.com/Phillezi/common/scripts/glint/internal/api"
	"github.com/Phillezi/common/scripts/glint/pkg/loader"
	"golang.org/x/tools/go/analysis"

	"github.com/gordonklaus/ineffassign/pkg/ineffassign"
)

type Command struct{}

func (c *Command) Run(path string, options ...api.CommandOption) error {
	opts := &api.Options{
		Logger: api.DefaultLogger,
	}
	for _, opt := range options {
		opt(opts)
	}

	opts.Logger.Infof("Running ineffassign on %s", path)

	prog, err := loader.Load(loader.Config{Dir: path})
	if err != nil {
		return fmt.Errorf("ineffassign: failed to load program: %w", err)
	}

	var issues []analysis.Diagnostic

	for _, pkg := range prog.Packages {
		pass := &analysis.Pass{
			Analyzer:  ineffassign.Analyzer,
			Fset:      prog.Fset,
			Files:     pkg.Syntax,
			Pkg:       pkg.Types,
			TypesInfo: pkg.TypesInfo,
			Report: func(d analysis.Diagnostic) {
				issues = append(issues, d)
			},
		}

		_, err := ineffassign.Analyzer.Run(pass)
		if err != nil {
			return fmt.Errorf("ineffassign analyzer error: %w", err)
		}
	}

	if len(issues) > 0 {
		opts.Logger.Errorf("ineffassign found unreachable or redundant assignments:")
		for _, issue := range issues {
			pos := prog.Fset.Position(issue.Pos)
			opts.Logger.PrintIndentf("%s:%d: %s", pos.Filename, pos.Line, issue.Message)
		}
		return fmt.Errorf("ineffassign: issues found")
	}

	opts.Logger.Successf("passed ineffassign")
	return nil
}
