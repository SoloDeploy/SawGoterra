package sawgoterra

func NewTerraformGraphParams() *TerraformParams {
	return &TerraformParams{
		Plan:        "terraform.tfplan",
		DrawCycles:  false,
		Type:        "plan",
		ModuleDepth: 0,
	}
}

func (t *TerraformCli) Graph(p *TerraformParams) error {
	return terraformAction("graph", t, p)
}
