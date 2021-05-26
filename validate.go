package sawgorerra

func NewTerraformValidateParams() *TerraformParams {
	return &TerraformParams{
		Json:    true,
		NoColor: true,
	}
}

func (t *TerraformCli) Validate(p *TerraformParams) error {
	return terraformAction("validate", t, p)
}

func (t *TerraformCli) ValidateWithDefaults() error {
	return t.Validate(NewTerraformValidateParams())
}
