package sawgoterra

func NewTerraformUntaintParams() *TerraformParams {
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

func (t *TerraformCli) Untaint(p *TerraformParams) error {
	return terraformAction("untaint", t, p)
}
