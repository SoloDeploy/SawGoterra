# sawgorerra
Terraform wrapper for Golang

## Example

```
package main

import (
	"fmt"
	"log"

	sg "github.com/SoloDeploy/sawgorerra"
)

func main() {

	tf, err := sg.NewTerraformCli()
	tf = tf.WithWorkingDirectory("src")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tf)

	initP := sg.NewTerraformInitParams()
	initP.BackendConfig = map[string]string{
		"subscriptionId": "abc",
	}

	err = tf.Init(initP)

	if err != nil {
		log.Fatal(err)
	}
}

```
