package sawgoterra

import (
	"testing"
)

func TestNewTerraformApplyParams(t *testing.T) {
	params := NewTerraformApplyParams()

	if params.Backup != "terraform.backup" {
		t.Errorf("Backup is not terraform.backup")
	}

	if params.AutoApprove != false {
		t.Errorf("AutoApprove is not false")
	}

	if params.Input != false {
		t.Errorf("Input is not false")
	}

	if params.Lock != true {
		t.Errorf("Lock is not true")
	}

	if params.LockTimeout != 0 {
		t.Errorf("LockTimeout is not 0")
	}

	if params.NoColor != true {
		t.Errorf("NoColor is not true")
	}

	if params.StateOut != "" {
		t.Errorf("StateOut is not empty")
	}

	if params.Parallelism != 10 {
		t.Errorf("Parallelism is not 10")
	}

	if params.State != "terraform.tfstate" {
		t.Errorf("State is not terraform.tfstate")
	}

	if params.Plan != "" {
		t.Errorf("StateOut is not empty")
	}
}
