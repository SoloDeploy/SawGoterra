package sawgoterra

func NewTerraformDestroyParams() *TerraformParams {
	return &TerraformParams{
		Backup:      "terraform.backup",
		AutoApprove: false,
		Lock:        true,
		LockTimeout: 0,
		NoColor:     true,
		Parallelism: 10,
		Refresh:     true,
		State:       "terraform.tfstate",
		StateOut:    "",
	}
}

func (t *TerraformCli) Destroy(p *TerraformParams) error {
	return terraformAction("destroy", t, p)
}
