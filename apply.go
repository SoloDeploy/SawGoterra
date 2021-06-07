package sawgoterra

func NewTerraformApplyParams() *TerraformParams {
	return &TerraformParams{
		Backup:      "terraform.backup",
		AutoApprove: false,
		Lock:        true,
		LockTimeout: 0,
		Input:       false,
		NoColor:     true,
		Parallelism: 10,
		Refresh:     true,
		State:       "terraform.tfstate",
		StateOut:    "",
		Plan:        "",
	}
}

func (t *TerraformCli) Apply(p *TerraformParams) error {
	return terraformAction("plan", t, p)
}
