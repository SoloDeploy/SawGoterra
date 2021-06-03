package sawgoterra

func NewTerraformImportParams() *TerraformParams {
	return &TerraformParams{
		Config:              "",
		AllowMissingConfig:  false,
		Input:               false,
		Lock:                true,
		LockTimeout:         0,
		NoColor:             true,
		Var:                 nil,
		VarFile:             nil,
		IgnoreRemoteVersion: false,
		State:               "",
		StateOut:            "",
		Backup:              "",
	}
}

func (t *TerraformCli) Import(p *TerraformParams) error {
	return terraformAction("import", t, p)
}
