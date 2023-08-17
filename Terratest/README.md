### Introduction

The Terratest Go library simplifies the process of writing automated  tests for your Terraform infrastructure. The following items can be  addressed with a single test using the Terratest library:

- Successful creation and deletion of a GCP resource
- Ensuring the correct input variables are provided and applied correctly
- Ensuring the resource module returns the expected amount of output variables

In this lab step, you will write a test to validate a Google Compute Engine instance module using the Terratest Go library. 

 

### Instructions

1. Expand the **terraform-gcp** directory below **PROJECT** to review the existing directory structure:

![alt](https://assets.cloudacademy.com/bakery/media/uploads/content_engine/image-20220117160539-1-8aab9b8e-3b11-4706-bb84-14a3237dbe8a.png)

In this lab, you will be testing a Google Compute Engine (GCE)  Terraform module. This directory structure may appear different compared to previous labs, as it utilizes the recommended testing structure  highlighted in the [Terratest Quickstart guide](https://terratest.gruntwork.io/docs/getting-started/quick-start/). 

The top-level **main.tf,** **outputs.tf**, and **variables.tf** files represent the single GCE server instance that will be deployed as a part of the test. The files of the same name within the **examples/gce** directory will be configured to run the Terraform process. 

 

2. Double-click on the **examples/gce/main.tf** file to open it in the editor.

 

3. Paste the following code into the **examples/gce/main.tf** file:

```
terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "3.5.0"
    }
  }
}
 
provider "google" {
  credentials = "/home/project/.sa_key"
  project = "cal-3032-160494a73309"
  region = "us-central1"
}
 
module "webserver" {
    source = "../../"
    name         = var.name
    machine_type = var.machine_type
    zone         = var.zone
}
```

You should be familiar with these configuration blocks. This will  configure the credentials and project settings for Terraform to access.  The `webserver module` defined at the bottom indicates it is  accessing the module stored at the root level of this directory and  passes in the three expected input variables. 

 

6. Double-click on the **examples/gce/outputs.tf** and **variables****.tf** files to review them:

![alt](https://assets.cloudacademy.com/bakery/media/uploads/content_engine/image-20220117161419-2-ca01401d-8de8-4b50-94a6-ba84d7c3d94a.png)

The **variables.tf** file establishes the three input variables for this **gce** module. They correspond with configurable settings of the GCE instance. 

![alt](https://assets.cloudacademy.com/bakery/media/uploads/content_engine/image-20220117161427-3-6d24f081-18b0-4666-b622-b83476708221.png)

The **outputs.tf** file establishes the two output variables for this **gce** module. Once it is deployed successfully, Terraform should output the IP address and the unique identifier of the GCE instance.

Both the input and output variables of this module will be tested to  ensure they are utilized correctly and returned as expected.

 

7. Double-click on the **terraform-gcp/main.tf** file at the root level of the project:

![alt](https://assets.cloudacademy.com/bakery/media/uploads/content_engine/image-20220117162103-4-ee2ab375-68f8-4166-b6e9-516c69ef2417.png)

To tie it all together, this is the GCE instance resource definition  that will be populated with the expected input variables. When the test  run begins, this definition will be used to create your Terraform  infrastructure. 

 

8. Double-click on the **test/gce_test.go** file to configure your test:

![alt](https://assets.cloudacademy.com/bakery/media/uploads/content_engine/image-20220117162323-5-ac1d7485-4e75-4e0b-9b52-cfcfb948b5e1.png)

This file will open as an empty file. 

 

9. Paste the following code into **test/gce_test.go**

```
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
```

At the top of the file, you can find the necessary imports for your test:

- `testing`: This import is necessary for implementing the automated test in Go, and will work in concert with the `go test` command you will run in the upcoming steps. 
- `terraform`: Imported directly from the GitHub repository, this import will allow  you to call Terraform actions throughout your test. Most notably, the  commands needed to set up, configure, and tear down your resource  module.
- `assert`: Imported directly from the GitHub  repository, this import provides access to assertions in Go. The various assertions accept input and will validate certain aspects of your  infrastructure by comparing actual and expected values. Assertions are  the building blocks of these types of Go tests and are very powerful.

Up next is the main test function `TestGCEWebserverCreate`. As you can tell by the name, this function will test for the successful creation of the GCE webserver instance. The creation and deletion  process of a Terraform resource will be referenced in this function in  order to test the complete process. 

The function's sole argument to observe is the `t *testing.T` object. This object will be passed in as an argument for each of the  Terraform methods called. It will act as a record keeper and will log  test results and failures, some of which will be output during the test  run. 

The first variable in the test function is the `terraformOptions` variable. It calls the `terraform.WithDefaultRetryableErrors` method to configure a Terraform project within the test. The two parameters are `TerraformDir`, and `Vars`. These parameters inform Terratest of the location of the Terraform  project and its expected input variables. This method will produce a  failed test result if the expected input variables aren't provided, or  if the expected variables are updated in some way. 

The next two methods, `InitAndApply`, and `Destroy` are called with the `terraformOptions` variable passed as an argument. The `InitAndApply` method will run the Terraform init and apply commands to spin up your Terraform infrastructure. The `Destroy` method is prepended with the `defer` key, which instructs Terratest to tear down the Terraform infrastructure *after* all the tests in this function have been completed. These methods will  produce a failed test result if the infrastructure encounters any errors during the creation or deletion process. 

Finally, the `outputMap` variable is declared by using the `OutputAll` method. This method will retrieve the entire output variable map object of your Terraform project. The `Equal` assertion will pass successfully if the total number of items in the `outputMap` variable is `2`. The built-in `len` method will return the number of items in the object passed in. You will recall that the two expected output variables are the `ip` and `instance_id` of the GCE server. 

With the existing resource configurations, you should expect this test to pass. 

 

10. In the terminal, change into the **terraform-gcp/test** directory with the following command:

```
cd terraform-gcp/test
```

You will initiate your test runs in this **test** directory.

 

11. In the terminal, run the following command to load the necessary test packages:

```
go mod tidy
```

You will notice a **go.sum** file appear in the **test** directory. The necessary packages have been downloaded and you can now run your test.

Within this same directory, the **go.mod** file will  also be updated. This file defines the required packages used both  directly and indirectly as well as the minimum required version of Go  needed to run these tests:

![alt](https://assets.cloudacademy.com/bakery/media/uploads/content_engine/image-20220117163328-6-f478ac95-425a-40d6-a763-b4873703d630.png)

These files have been pre-configured for you, so don't worry too much about their configurations. 

 

12. In the terminal, run the following command to run your test:

```
go test -v gce_test.go
```

The `-v` option provides additional output during the test run:

![alt](https://assets.cloudacademy.com/bakery/media/uploads/content_engine/image-20220117164109-7-fe2282ee-9d8a-428f-992b-040638477e1b.png)

![alt](https://assets.cloudacademy.com/bakery/media/uploads/content_engine/image-20220117164121-8-62ccf51a-7387-4c1c-a613-b4be35268416.png)

The test begins with a Terraform initialization, followed by the creation plan output.

![alt](https://assets.cloudacademy.com/bakery/media/uploads/content_engine/image-20220117164230-9-a2318ecb-2322-44ac-987d-c45bc80e06fe.png)

![alt](https://assets.cloudacademy.com/bakery/media/uploads/content_engine/image-20220117164255-10-d61f9ec5-47e7-437f-8f26-2dec7c0acb7f.png)

The resource is deployed and you will notice the output variable  object being printed to the terminal. It is at this point that Terratest will initiate your tests with the test results saved for final output.

![alt](https://assets.cloudacademy.com/bakery/media/uploads/content_engine/image-20220117165154-11-85f1d089-2098-4ce9-88bb-13916dc141fa.png)

![alt](https://assets.cloudacademy.com/bakery/media/uploads/content_engine/image-20220117165215-12-f1a8189d-265d-420a-bc7e-df8adbdbf497.png)

Finally, your GCE instance is destroyed after the tests have been  completed, and the final test result is printed on the screen. The **TestGCEWebserverCreate** test has successfully passed!

 

### Summary

In this lab, you wrote a test to validate a Google Compute Engine  instance module using the Terratest Go library. You also learned about  the inner workings of the Terratest library and how to structure a test  directory. 
