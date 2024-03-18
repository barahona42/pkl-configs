package pklconfigs_test

import (
	"context"
	"testing"

	"github.com/barahona42/pklconfigs/pkg/applications/directionsapp"
)

func TestMain(t *testing.T) {
	ac, err := directionsapp.LoadFromPath(context.TODO(), ".local/config.pkl")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ac.Application.Name)
	for i := range ac.Configuration.Routes {
		t.Log(ac.Configuration.Routes[i].Label)
	}
}
