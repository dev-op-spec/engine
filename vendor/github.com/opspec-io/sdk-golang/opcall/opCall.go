// Package opcall implements usecases surrounding op calls
package opcall

//go:generate counterfeiter -o ./fake.go --fake-name Fake ./ OpCall

import (
	"github.com/opspec-io/sdk-golang/expression"
	"github.com/opspec-io/sdk-golang/model"
	"github.com/opspec-io/sdk-golang/opcall/inputs"
	"github.com/opspec-io/sdk-golang/pkg"
	"github.com/opspec-io/sdk-golang/util/uniquestring"
	"path/filepath"
)

type OpCall interface {
	// Interpret interprets an SCGOpCall into a DCGOpCall
	Interpret(
		scope map[string]*model.Value,
		scgOpCall *model.SCGOpCall,
		opId string,
		parentPkgHandle model.PkgHandle,
		rootOpId string,
	) (*model.DCGOpCall, error)
}

func New(
	rootFSPath string,
) OpCall {
	return _OpCall{
		dcgScratchDir:       filepath.Join(rootFSPath, "dcg"),
		expression:          expression.New(),
		pkg:                 pkg.New(),
		pkgCachePath:        filepath.Join(rootFSPath, "pkgs"),
		uniqueStringFactory: uniquestring.NewUniqueStringFactory(),
		inputs:              inputs.New(),
	}
}

type _OpCall struct {
	dcgScratchDir       string
	expression          expression.Expression
	pkg                 pkg.Pkg
	pkgCachePath        string
	uniqueStringFactory uniquestring.UniqueStringFactory
	inputs              inputs.Inputs
}
