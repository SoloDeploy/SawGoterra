package sawgorerra

type TerraformDestoryParams struct {
	Backup      string
	AutoApprove bool
	Lock        bool
	LockTimeout int
	NoColor     bool
	Parallelism int
	Refresh     bool
	State       string
	StateOut    string
}

func NewTerraformDestoryParams() *TerraformDestoryParams {
	return &TerraformDestoryParams{
		Backup:      "terraform.backup",
		AutoApprove: false,
		Lock:        true,
		LockTimeout: 0,
		NoColor:     true,
		Parallelism: 10,
		Refresh:     true,
		State:       "terraform.tfstate",
		StateOut:    "",
	}
}
