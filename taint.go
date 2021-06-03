package sawgoterra

func NewTerraformTaintParams() *TerraformParams {
	return &TerraformParams{
		AllowMissing:        false,
		Backup:              "",
		Lock:                true,
		LockTimeout:         0,
		State:               "terraform.tfstate",
		StateOut:            "",
		IgnoreRemoteVersion: false,
	}
}

func (t *TerraformCli) Taint(p *TerraformParams) error {
	return terraformAction("taint", t, p)
}
