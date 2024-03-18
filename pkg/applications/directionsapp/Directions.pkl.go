// Code generated from Pkl module `barahona42.configs.pkg.applications.DirectionsApp`. DO NOT EDIT.
package directionsapp

import "github.com/apple/pkl-go/pkl"

type Directions struct {
	PollRate *pkl.Duration `pkl:"PollRate"`

	RunTime *pkl.Duration `pkl:"RunTime"`

	ApiKey string `pkl:"ApiKey"`

	Routes []*Route `pkl:"Routes"`

	RoutesFile string `pkl:"RoutesFile"`

	Routesdb string `pkl:"Routesdb"`
}
