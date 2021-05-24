package sawgorerra

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"regexp"
	"strings"
)

func terraformAction(a string, t *TerraformCli, p interface{}) error {
	tpl, err := template.New(a+".tmpl").ParseFS(templates, "partials/*.tmpl", "templates/"+a+".tmpl")

	if err != nil {
		return err
	}

	var cmd bytes.Buffer
	type Data struct {
		Cli    *TerraformCli
		Params interface{}
	}
	data := &Data{t, p}
	err = tpl.Execute(&cmd, data)

	if err != nil {
		log.Fatal(err)
	}

	space := regexp.MustCompile(`\s+`)
	s := space.ReplaceAllString(cmd.String(), " ")
	cmdStr := strings.Trim(s, " ")
	fmt.Println(cmdStr)

	fmt.Println(t.Path, strings.Split(cmdStr, " "))
	tfCmd := execCommand(t.Path, strings.Split(cmdStr, " ")...)
	out, err := tfCmd.Output()

	if err != nil {
		return err
	}

	fmt.Println(string(out))

	return nil
}
