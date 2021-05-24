package sawgorerra

import (
	"embed"
	"os/exec"
	"strings"
)

var execCommand = exec.Command

//go:embed partials/*.tmpl templates/*.tmpl
var templates embed.FS

type TerraformCli struct {
	Path       string
	Version    string
	WorkingDir string
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
	out, err := execCommand(t.Path, "--version").Output()

	if err != nil {
		return nil, err
	}

	outStr := strings.Split(string(out), "\n")
	t.Version = strings.Split(outStr[0], " v")[1]
	return t, nil
}
