package sawgorerra

type TerraformValidateParams struct {
	Json    bool
	NoColor bool
}

func NewTerraformValidateParams() *TerraformValidateParams {
	return &TerraformValidateParams{
		Json:    true,
		NoColor: true,
	}
}
