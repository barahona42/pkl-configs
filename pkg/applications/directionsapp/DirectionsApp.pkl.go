// Code generated from Pkl module `barahona42.configs.applications.DirectionsApp`. DO NOT EDIT.
package directionsapp

import (
	"context"

	"configs/pkg/applications/core"
	"github.com/apple/pkl-go/pkl"
)

type DirectionsApp struct {
	Application *core.Application `pkl:"Application"`

	Configuration *Directions `pkl:"Configuration"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a DirectionsApp
func LoadFromPath(ctx context.Context, path string) (ret *DirectionsApp, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a DirectionsApp
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*DirectionsApp, error) {
	var ret DirectionsApp
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
