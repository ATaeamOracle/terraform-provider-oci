    #     ___  ____     _    ____ _     _____
    #    / _ \|  _ \   / \  / ___| |   | ____|
    #   | | | | |_) | / _ \| |   | |   |  _|  
    #   | |_| |  _ < / ___ | |___| |___| |___
    #    \___/|_| \_/_/   \_\____|_____|_____|
***
# Terraform provider for Oracle Bare Metal Cloud Services
Oracle customers now have access to an enterprise class, developer friendly orchestration tool they can use to manage [Oracle Bare Metal Cloud Service](https://cloud.oracle.com/en_US/bare-metal) resources as well as the [Oracle Compute Cloud](https://github.com/oracle/terraform-provider-compute).

This Terraform provider is OSS, available to all OBMCS customers at no charge.

### Coverage
The Terraform provider provides coverage for the entire BMC API excluding the Load Balancer Service, expected March 1 2017.  

## Getting started
Be sure to read the FAQ and Writing Terraform configurations for OBMCS in [/docs](https://github.com/oracle/terraform-provider-baremetal/docs).

### Install Terraform
https://www.terraform.io/downloads.html

### Download the OBMCS Terraform provider binary
Find the appropriate binary for [your platform here](https://github.com/oracle/terraform-provider-baremetal/releases), download it.

#### \*nix
Create `~/.terraformrc` that specifies the path to the `baremetal` provider.  
```
providers {
  baremetal = "<path_to_provider_binary>"
  }
```

#### Windows
Create `%APPDATA%/terraform.rc` that specifies the path to the `baremetal` provider.
```
providers {
  baremetal = "<path_to_provider_binary>"
  }
```
### Export credentials
Required Keys and OCIDs - https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm  

If you primarily work in a single compartment consider exporting that compartment's OCID as well. Remember that the tenancy OCID is also the OCID of the root compartment.

#### \*nix
In your ~/.bash_profile set these variables  
`export TF_VAR_tenancy_ocid=`  
`export TF_VAR_user_ocid=`  
`export TF_VAR_fingerprint=`  
`export TF_VAR_private_key_path=<fully qualified path>`  
`export TF_VAR_private_key_password=`  

Don't forget to `source ~/.bash_profile` once you've set these.

#### Windows
`setx TF_VAR_tenancy_ocid <value>`  
`setx TF_VAR_user_ocid <value>`  
`setx TF_VAR_fingerprint <value>`  
`setx TF_VAR_private_key_path <value>`  
`setx TF_VAR_private_key_password <value>`  

The variables won't be set for the current session, exit the terminal and reopen.

## Deploy an example configuration
Download the [VCN example configuration.](https://github.com/oracle/terraform-provider-baremetal/docs/examples/simple_vcn)  

Edit it to include the OCID of the compartment you want to create the VCN. Remember that the tenancy OCID is the compartment OCID of your root compartment.

You should always plan, then apply a configuration -  
```
$ terraform plan VCN.tf
# Make sure the plan looks right.
$ terraform apply VCN.tf
```
## OBMC resource and datasource details
https://github.com/oracle/terraform-provider-baremetal/tree/master/docs

## Getting help
You can file an issue against the project  
https://github.com/oracle/terraform-provider-baremetal/issues

or meet us in the OBMCS forums  
https://community.oracle.com/community/cloud_computing/bare-metal

## Known bugs
1. The private key you use for API access must have a passphrase to work with Terraform. You can add a passphrase to your existing key - `ssh-keygen -p -f <private key>`

2. The provider doesn't always wait for resources to transition to 'AVAILABLE' before exiting and will throw an error if any resource isn't 'AVAILABLE' when the configuration has finished being applied. You can safely ignore these errors.

#### About the provider
This provider was written on behalf of Oracle by [MustWin.](http://mustwin.com/)
