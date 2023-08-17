package test
 
import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)
 
func TestGCEWebserverCreate(t *testing.T) {
 
	// The values to pass into the Terraform CLI
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		  
        // The path to where the example Terraform code is located
		TerraformDir: "../examples/gce",
		// Variables to pass to the Terraform code using -var options
		Vars: map[string]interface{}{
			"name":    "server",
			"machine_type": "f1-micro",
			"zone": "us-central1-a",
		},
	})
   
	// Run a Terraform init and apply with the Terraform options
	terraform.InitAndApply(t, terraformOptions)
  
	// Run a Terraform Destroy at the end of the test
	defer terraform.Destroy(t, terraformOptions)
  
	// Validate Terraform Output contains 2 values
	outputMap := terraform.OutputAll(t, terraformOptions)
	assert.Equal(t, 2, len(outputMap))
}