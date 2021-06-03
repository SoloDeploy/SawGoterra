package sawgoterra

func NewTerraformProvidersParams() *TerraformParams {
	return &TerraformParams{
		Json:    true,
		Raw:     false,
		NoColor: true,
		State:   "terraform.tfstate",
	}
}

func (t *TerraformCli) Providers(p *TerraformParams) error {
	return terraformAction("providers", t, p)
}
