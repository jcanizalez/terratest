package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
	 http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
     "time"
)

func TestTerraformFullWebApp(t *testing.T) {


	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../terraform-azure-full-web-app",
	
	})

	// Clean up resources with "terraform destroy" at the end of the test.
	defer terraform.Destroy(t, terraformOptions)

	// Run "terraform init" and "terraform apply". Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables and check they have the expected values.
	webUrl := terraform.Output(t, terraformOptions, "web_url")
	http_helper.HttpGetWithRetryWithCustomValidation(t, webUrl, nil, 30, 5*time.Second,func(status int,body string) bool { return status == 200})
}