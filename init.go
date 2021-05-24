package sawgorerra

type TerraformInitParams struct {
	Backend       bool
	BackendConfig map[string]string
	ForceCopy     bool
	FromModule    string
	Get           bool
	Input         bool
	NoColor       bool
	PluginDir     []string
	Reconfigure   bool
	Upgrade       bool
	Lockfile      string
}

func NewTerraformInitParams() *TerraformInitParams {
	return &TerraformInitParams{
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

func (t *TerraformCli) Init(p *TerraformInitParams) error {
	err := terraformAction("init", t, p)
	return err
}
