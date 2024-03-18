package configs_test

import (
	"configs/pkg/applications/directionsapp"
	"context"
	"testing"
)

func TestMain(t *testing.T) {
	ac, err := directionsapp.LoadFromPath(context.TODO(), ".local/config.pkl")
	// ac, err := appconfig.LoadFromPath(context.TODO(), ".local/config.pkl")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ac.Application.Name)
	for i := range ac.Configuration.Routes {
		t.Log(ac.Configuration.Routes[i].Label)
	}
}
