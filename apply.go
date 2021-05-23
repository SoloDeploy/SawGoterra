package sawgorerra

type TerraformApplyParams struct {
	Backup      string
	AutoApprove bool
	Lock        bool
	LockTimeout int
	Input       bool
	NoColor     bool
	Parallelism int
	Refresh     bool
	State       string
	StateOut    string
	Plan        string
}

func NewTerraformApplyParams() *TerraformApplyParams {
	return &TerraformApplyParams{
		Backup:      "terraform.backup",
		AutoApprove: false,
		Lock:        true,
		LockTimeout: 0,
		Input:       false,
		NoColor:     true,
		Parallelism: 10,
		Refresh:     true,
		State:       "terraform.tfstate",
		StateOut:    "",
		Plan:        "",
	}
}
