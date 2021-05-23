package sawgorerra

import (
	"bytes"
	"fmt"
	"log"
	"regexp"
	"text/template"
)

type TerraformInitParams struct {
	Backend       bool
	BackendConfig map[string]string
	ForceCopy     bool
	FromModule    string
	Get           bool
	Input         bool
	NoColor       bool
	PluginDir     []string
	Reconfigure   bool
	Upgrade       bool
	Lockfile      string
}

func NewTerraformInitParams() *TerraformInitParams {
	return &TerraformInitParams{
		Backend:       true,
		BackendConfig: nil,
		ForceCopy:     false,
		FromModule:    "",
		Get:           false,
		Input:         true,
		NoColor:       true,
		PluginDir:     nil,
		Reconfigure:   false,
		Upgrade:       false,
		Lockfile:      "readonly",
	}
}

func (t *TerraformCli) Init(p *TerraformInitParams) error {
	fmt.Println(t)
	tpl, err := template.New("init.tmpl").ParseFS(templates, "partials/*.tmpl", "templates/init.tmpl")

	if err != nil {
		log.Fatal(err)
	}

	var cmd bytes.Buffer
	type Data struct {
		Cli    *TerraformCli
		Params *TerraformInitParams
	}
	data := &Data{t, p}
	err = tpl.Execute(&cmd, data)

	if err != nil {
		log.Fatal(err)
	}

	space := regexp.MustCompile(`\s+`)
	s := space.ReplaceAllString(cmd.String(), " ")

	fmt.Println(s)
	fmt.Println("done")

	return nil
}
