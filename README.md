# sawgoterra

Terraform wrapper for Golang

## Example

```go
package main

import (
  "log"

  sg "github.com/SoloDeploy/sawgoterra"
)

var backendConfig = map[string]string{
  "subscription_id": "abc"
}

var defaultVars = map[string]string{
  "location": "uksouth"
}

func main() {

  tf, err := sg.NewTerraformCli()
  tf = tf.WithWorkingDirectory("src")

  if err != nil {
    log.Fatal(err)
  }

  err = tf.InitWithBackendConfig(backendConfig)
  if err != nil {
    log.Fatal(err)
  }

  err = tf.ValidateWithDefaults()
  if err != nil {
    log.Fatal(err)
  }

  planP = sg.NewTerraformPlanParams()
  planP.Var = defaultVars
  planP.Out = "terraform.tfplan"
  err = tf.Plan(planP)
  if err != nil {
    log.Fatal(err)
  }

  applyP = sg.NewTerraformApplyParams()
  applyP.Plan = "terraform.tfplan"
  err = tf.Apply(applyP)
  if err != nil {
    log.Fatal(err)
  }
}
```
