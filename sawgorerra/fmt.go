package sawgorerra

func NewTerraformFmtParams() *TerraformParams {
	return &TerraformParams{
		List:      true,
		Write:     true,
		Diff:      false,
		Check:     false,
		NoColor:   true,
		Recursive: false,
	}
}

func (t *TerraformCli) Fmt(p *TerraformParams) error {
	return terraformAction("fmt", t, p)
}
