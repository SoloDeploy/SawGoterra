package sawgorerra

func NewTerraformGetParams() *TerraformParams {
	return &TerraformParams{
		Update:  true,
		NoColor: true,
	}
}

func (t *TerraformCli) Get(p *TerraformParams) error {
	return terraformAction("get", t, p)
}
