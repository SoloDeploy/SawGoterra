package sawgoterra

func NewTerraformRefreshParams() *TerraformParams {
	return &TerraformParams{
		Backup:          "",
		CompactWarnings: false,
		Input:           false,
		Lock:            true,
		LockTimeout:     0,
		NoColor:         true,
		Target:          nil,
		Var:             nil,
		VarFile:         nil,
		State:           "",
		StateOut:        "",
	}
}

func (t *TerraformCli) Refresh(p *TerraformParams) error {
	return terraformAction("refresh", t, p)
}
