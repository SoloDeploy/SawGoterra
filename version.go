package sawgoterra

func (t *TerraformCli) Version() error {
	return terraformAction("version", t, &TerraformParams{})
}
