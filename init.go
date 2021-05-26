package sawgorerra

func NewTerraformInitParams() *TerraformParams {
	return &TerraformParams{
		Backend:       true,
		BackendConfig: nil,
		ForceCopy:     false,
		FromModule:    "",
		Get:           false,
		Input:         true,
		NoColor:       true,
		PluginDir:     nil,
		Reconfigure:   false,
		Upgrade:       false,
		Lockfile:      "",
	}
}

func (t *TerraformCli) Init(p *TerraformParams) error {
	err := terraformAction("init", t, p)
	return err
}

func (t *TerraformCli) InitWithBackendConfig(backendConfig map[string]string) error {
	p := NewTerraformInitParams()
	p.BackendConfig = backendConfig
	return t.Init(p)
}
