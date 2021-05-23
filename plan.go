package sawgorerra

type TerraformPlanParams struct {
	Destroy          bool
	Refresh          bool
	Replace          string
	Target           []string
	Var              []string
	VarFile          []string
	CompactWarnings  bool
	DetailedExitcode bool
	Input            bool
	Lock             bool
	LockTimeout      int
	NoColor          bool
	Out              string
	Parallelism      int
	State            string
}

func NewTerraformPlanParams() *TerraformPlanParams {
	return &TerraformPlanParams{
		Destroy:          false,
		Refresh:          false,
		Replace:          "",
		Target:           nil,
		Var:              nil,
		VarFile:          nil,
		CompactWarnings:  false,
		DetailedExitcode: false,
		Input:            false,
		Lock:             true,
		LockTimeout:      0,
		NoColor:          true,
		Out:              "",
		Parallelism:      10,
		State:            "terraform.tfstate",
	}
}
