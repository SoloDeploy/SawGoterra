package sawgoterra

func NewTerraformPlanParams() *TerraformParams {
	return &TerraformParams{
		Destroy:          false,
		Refresh:          false,
		Replace:          "",
		Target:           nil,
		Var:              nil,
		VarFile:          nil,
		CompactWarnings:  false,
		DetailedExitcode: false,
		Input:            false,
		Lock:             true,
		LockTimeout:      0,
		NoColor:          true,
		Out:              "",
		Parallelism:      10,
		State:            "terraform.tfstate",
	}
}

func (t *TerraformCli) Plan(p *TerraformParams) error {
	return terraformAction("plan", t, p)
}
