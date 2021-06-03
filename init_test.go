package sawgoterra

import (
	"testing"
)

func TestNewTerraformInitParams(t *testing.T) {
	params := NewTerraformInitParams()

	if params.Backend != true {
		t.Errorf("Backend is not terraform")
	}

	if params.BackendConfig != nil {
		t.Errorf("BackendConfig is not nil")
	}

	if params.ForceCopy != false {
		t.Errorf("ForceCopy is not false")
	}

	if params.FromModule != "" {
		t.Errorf("FromModule is not empty")
	}

	if params.Get != false {
		t.Errorf("Get is not false")
	}

	if params.Input != true {
		t.Errorf("Input is not true")
	}

	if params.NoColor != true {
		t.Errorf("NoColor is not true")
	}

	if params.PluginDir != nil {
		t.Errorf("PluginDir is not nil")
	}

	if params.Reconfigure != false {
		t.Errorf("Reconfigure is not false")
	}

	if params.Upgrade != false {
		t.Errorf("Upgrade is not false")
	}

	if params.Lockfile != "" {
		t.Errorf("Lockfile is not empty")
	}
}
