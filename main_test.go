package sawgoterra

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

var tfVersionStr = `Terraform v0.11.0
on darwin_amd64
`

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	// some code here to check arguments perhaps?
	fmt.Fprintf(os.Stdout, tfVersionStr)
	os.Exit(0)
}

func TestNewTerraformCliWithPath(t *testing.T) {
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()
	cli, _ := NewTerraformCliWithPath("bin/terraform0.11.0")

	if cli.Path != "bin/terraform0.11.0" {
		t.Errorf("Path is not terraform")
	}

	if cli.VersionNumber != "0.11.0" {
		t.Errorf("VersionNumber is incorrect")
	}
}
