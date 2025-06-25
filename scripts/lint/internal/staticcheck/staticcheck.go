package staticcheck

import (
	"fmt"
	"go/token"

	"github.com/Phillezi/common/scripts/glint/internal/api"
	"github.com/Phillezi/common/scripts/glint/pkg/loader"
	"golang.org/x/tools/go/analysis"
	"honnef.co/go/tools/staticcheck"
)

type Command struct{}

func (c *Command) Run(path string, options ...api.CommandOption) error {
	opts := &api.Options{
		Logger: api.DefaultLogger,
	}
	for _, opt := range options {
		opt(opts)
	}

	opts.Logger.Infof("Running staticcheck on %s", path)

	cfg := loader.Config{
		Dir: path,
	}
	program, err := loader.Load(cfg)
	if err != nil {
		return fmt.Errorf("staticcheck: %w", err)
	}

	checks := staticcheck.Analyzers
	fset := token.NewFileSet()
	var found bool

	for _, a := range checks {
		pass := loader.CreatePass(a.Analyzer, fset, program)
		results, err := a.Analyzer.Run(pass)
		if err != nil {
			continue // skip errored checks
		}
		switch res := results.(type) {
		case []analysis.Diagnostic:
			for _, diag := range res {
				if !found {
					opts.Logger.Errorf("staticcheck reported issues:")
					found = true
				}
				pos := fset.Position(diag.Pos)
				opts.Logger.PrintIndentf("[%s] %s:%d: %s", a.Analyzer.Name, pos.Filename, pos.Line, diag.Message)
			}
		default:
			opts.Logger.Errorf("Unexpected type of results: %T", res)
		}
	}

	if found {
		return fmt.Errorf("staticcheck: issues found")
	}
	opts.Logger.Successf("passed staticcheck")
	return nil
}
