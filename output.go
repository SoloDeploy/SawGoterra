package sawgoterra

func NewTerraformOutputParams() *TerraformParams {
	return &TerraformParams{
		Json:    true,
		Raw:     false,
		NoColor: true,
		State:   "terraform.tfstate",
	}
}

func (t *TerraformCli) Output(p *TerraformParams) error {
	return terraformAction("output", t, p)
}
