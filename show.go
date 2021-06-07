package sawgoterra

func NewTerraformShowParams() *TerraformParams {
	return &TerraformParams{
		NoColor: true,
		Json:    true,
	}
}

func (t *TerraformCli) Show(p *TerraformParams) error {
	return terraformAction("show", t, p)
}
