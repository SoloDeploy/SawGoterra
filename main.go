package sawgorerra

import (
	"embed"
	"fmt"
	"os/exec"
	"strings"
)

var execCommand = exec.Command

//go:embed partials/*.tmpl templates/*.tmpl
var templates embed.FS

type TerraformCli struct {
	Path          string
	VersionNumber string
	WorkingDir    string
}

type TerraformParams struct {
	AllowMissing        bool
	AllowMissingConfig  bool
	AutoApprove         bool
	Backend             bool
	BackendConfig       map[string]string
	Backup              string
	Check               bool
	CompactWarnings     bool
	Config              string
	Destroy             bool
	DetailedExitcode    bool
	Diff                bool
	DrawCycles          bool
	ForceCopy           bool
	FromModule          string
	Get                 bool
	IgnoreRemoteVersion bool
	Input               bool
	Json                bool
	List                bool
	Lock                bool
	Lockfile            string
	LockTimeout         int
	ModuleDepth         int
	NoColor             bool
	Out                 string
	Parallelism         int
	Plan                string
	PluginDir           []string
	Raw                 bool
	Reconfigure         bool
	Recursive           bool
	Refresh             bool
	Replace             string
	State               string
	StateOut            string
	Target              []string
	Type                string
	Update              bool
	Upgrade             bool
	Var                 map[string]string
	VarFile             []string
	Write               bool
}

func NewTerraformCli() (*TerraformCli, error) {
	return NewTerraformCliWithPath("terraform")
}

func NewTerraformCliWithPath(binPath string) (*TerraformCli, error) {
	cli := &TerraformCli{
		Path: binPath,
	}
	cli, err := cli.fetchVersion()
	return cli, err
}

func (t *TerraformCli) WithWorkingDirectory(workingDir string) *TerraformCli {
	t.WorkingDir = workingDir
	return t
}

func (t *TerraformCli) fetchVersion() (*TerraformCli, error) {
	fmt.Println("fetchVersion")
	out, err := execCommand(t.Path, "--version").Output()
	fmt.Println(err)
	fmt.Println(string(out))

	if err != nil {
		return nil, err
	}

	outStr := strings.Split(string(out), "\n")
	t.VersionNumber = strings.Split(outStr[0], " v")[1]
	return t, nil
}
