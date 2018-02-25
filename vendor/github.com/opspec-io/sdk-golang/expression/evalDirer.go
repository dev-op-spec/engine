package expression

import (
	"fmt"
	"github.com/opspec-io/sdk-golang/data"
	"github.com/opspec-io/sdk-golang/expression/interpolater"
	"github.com/opspec-io/sdk-golang/model"
	"path/filepath"
	"strings"
)

type evalDirer interface {
	// EvalToDir evaluates an expression to a dir value
	//
	// Examples of valid dir expressions:
	// scope ref: $(scope-ref)
	// scope ref w/ path: $(scope-ref/sub-dir)
	// scope ref w/ deprecated path: $(scope-ref)/sub-dir
	// pkg fs ref: $(/pkg-fs-ref)
	// pkg fs ref w/ path: $(/pkg-fs-ref/sub-dir)
	EvalToDir(
		scope map[string]*model.Value,
		expression string,
		pkgHandle model.PkgHandle,
	) (*model.Value, error)
}

func newEvalDirer() evalDirer {
	return _evalDirer{
		data:         data.New(),
		interpolater: interpolater.New(),
	}
}

type _evalDirer struct {
	data         data.Data
	interpolater interpolater.Interpolater
}

func (ed _evalDirer) EvalToDir(
	scope map[string]*model.Value,
	expression string,
	pkgHandle model.PkgHandle,
) (*model.Value, error) {
	possibleRefCloserIndex := strings.Index(expression, interpolater.RefEnd)

	// the following is gross but it's due to all the deprecated syntax we support
	if ref, ok := tryResolveExplicitRef(expression, scope); ok && nil != ref.Dir {
		// scope ref w/out path
		return ref, nil
	} else if strings.HasPrefix(expression, "/") {

		// deprecated pkg fs ref
		deprecatedPkgFsRefPath, err := ed.interpolater.Interpolate(
			expression,
			scope,
			pkgHandle,
		)
		if nil != err {
			return nil, fmt.Errorf("unable to evaluate %v to dir; error was %v", expression, err.Error())
		}

		deprecatedPkgFsRefPath = filepath.Join(pkgHandle.Ref(), deprecatedPkgFsRefPath)
		return &model.Value{Dir: &deprecatedPkgFsRefPath}, err

	} else if strings.HasPrefix(expression, interpolater.RefStart) && possibleRefCloserIndex > 0 {

		refExpression := expression[2:possibleRefCloserIndex]
		refParts := strings.SplitN(refExpression, "/", 2)
		var dirValue string

		if strings.HasPrefix(refExpression, "/") {

			// pkg fs ref
			pkgFsRef, err := ed.interpolater.Interpolate(refExpression, scope, pkgHandle)
			if nil != err {
				return nil, fmt.Errorf("unable to evaluate pkg fs ref %v; error was %v", refExpression, err.Error())
			}
			pkgPath := pkgHandle.Path()
			dirValue = filepath.Join(*pkgPath, pkgFsRef)

		} else if dcgValue, ok := scope[refExpression]; ok && nil != dcgValue.Dir {

			// dir scope ref w/ deprecated path
			deprecatedPathExpression := expression[possibleRefCloserIndex+1:]
			deprecatedPath, err := ed.interpolater.Interpolate(deprecatedPathExpression, scope, pkgHandle)
			if nil != err {
				return nil, fmt.Errorf("unable to evaluate path %v; error was %v", deprecatedPathExpression, err.Error())
			}

			dirValue := filepath.Join(*dcgValue.Dir, deprecatedPath)
			return &model.Value{Dir: &dirValue}, nil

		} else if dcgValue, ok := scope[refParts[0]]; ok && nil != dcgValue.Dir {

			// scope ref w/ path
			pathExpression := refParts[1]
			path, err := ed.interpolater.Interpolate(pathExpression, scope, pkgHandle)
			if nil != err {
				return nil, fmt.Errorf("unable to evaluate path %v; error was %v", pathExpression, err.Error())
			}

			// no possibility of deprecated path due to presence of path
			dirValue := filepath.Join(*dcgValue.Dir, path)
			return &model.Value{Dir: &dirValue}, nil

		}

		if len(expression) > possibleRefCloserIndex+1 {
			// deprecated path
			deprecatedPathExpression := expression[possibleRefCloserIndex+1:]
			deprecatedPath, err := ed.interpolater.Interpolate(deprecatedPathExpression, scope, pkgHandle)
			if nil != err {
				return nil, fmt.Errorf("unable to evaluate path %v; error was %v", deprecatedPathExpression, err.Error())
			}

			dirValue := filepath.Join(dirValue, deprecatedPath)
			return &model.Value{Dir: &dirValue}, nil
		}

		return &model.Value{Dir: &dirValue}, nil
	}

	return nil, fmt.Errorf("unable to evaluate %v to dir", expression)

}
