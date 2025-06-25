package loader

import (
	"context"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/packages"
)

type Config struct {
	Dir string
}

type Program struct {
	Fset     *token.FileSet
	Packages []*packages.Package
}

func Load(cfg Config) (*Program, error) {
	fset := token.NewFileSet()
	pkgs, err := packages.Load(&packages.Config{
		Context:    context.Background(),
		Dir:        cfg.Dir,
		Mode:       packages.LoadAllSyntax,
		Fset:       fset,
		Tests:      false,
		BuildFlags: []string{}, // set build tags or flags if needed
	}, "./...")
	if err != nil {
		return nil, err
	}
	return &Program{
		Fset:     fset,
		Packages: pkgs,
	}, nil
}

func CreatePass(a *analysis.Analyzer, fset *token.FileSet, prog *Program) *analysis.Pass {
	// NOTE: This is a *minimal* pass; for full feature support, more fields may need to be filled.
	return &analysis.Pass{
		Analyzer:  a,
		Fset:      fset,
		Files:     prog.Packages[0].Syntax,
		Pkg:       prog.Packages[0].Types,
		TypesInfo: prog.Packages[0].TypesInfo,
		ResultOf:  make(map[*analysis.Analyzer]interface{}),
		Report: func(d analysis.Diagnostic) {
			// Store it or print it externally
		},
	}
}
