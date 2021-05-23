package sawgorerra

type TerraformFmtParams struct {
	List bool
	Write bool
	Diff bool
	Check bool
	NoColor bool
	Recursive bool
}

func NewTerraformFmtParams() *TerraformFmtParams {
	return &TerraformFmtParams{
		List: true,
		Write: true,
		Diff: false,
		Check: false,
		NoColor: true,
		Recursive: false,
	}
}
