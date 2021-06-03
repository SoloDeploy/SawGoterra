package sawgoterra

import (
	"bytes"
	"fmt"
	"log"
	"regexp"
	"strings"
	"text/template"
)

func terraformAction(a string, t *TerraformCli, p *TerraformParams) error {
	tpl, err := template.New(a+".tmpl").ParseFS(templates, "partials/*.tmpl", "templates/"+a+".tmpl")

	if err != nil {
		return err
	}

	var cmd bytes.Buffer

	type Data struct {
		Cli    *TerraformCli
		Params *TerraformParams
	}

	data := &Data{t, p}
	err = tpl.Execute(&cmd, data)

	if err != nil {
		log.Fatal(err)
	}

	space := regexp.MustCompile(`\s+`)
	s := space.ReplaceAllString(cmd.String(), " ")
	cmdStr := strings.Trim(s, " ")

	tfCmd := execCommand(t.Path, strings.Split(cmdStr, " ")...)
	out, err := tfCmd.CombinedOutput()

	if err != nil {
		return err
	}

	fmt.Println(string(out))

	return nil
}
