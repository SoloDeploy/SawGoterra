package sawgoterra

import (
	"testing"
)

func TestNewTerraformPlanParams(t *testing.T) {
	params := NewTerraformPlanParams()

	if params.Destroy != false {
		t.Errorf("Destroy is not false")
	}

	if params.Refresh != false {
		t.Errorf("Refres is not false")
	}

	if params.Replace != "" {
		t.Errorf("Replace is not empty")
	}

	if params.Target != nil {
		t.Errorf("Target is not empty")
	}

	if params.Var != nil {
		t.Errorf("Var is not empty")
	}

	if params.VarFile != nil {
		t.Errorf("VarFile is not empty")
	}

	if params.CompactWarnings != false {
		t.Errorf("CompactWarnings is not false")
	}

	if params.DetailedExitcode != false {
		t.Errorf("DetailedExitcode is not false")
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

	if params.Out != "" {
		t.Errorf("Out is not empty")
	}

	if params.Parallelism != 10 {
		t.Errorf("Parallelism is not 10")
	}

	if params.State != "terraform.tfstate" {
		t.Errorf("State is not terraform.tfstate")
	}
}
