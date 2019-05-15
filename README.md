# terratest-sample

Sample extracted from the doc [Test Terraform modules in Azure by using Terratest](https://docs.microsoft.com/en-us/azure/terraform/terratest-in-terraform-modules)

## Requirements

- **Go programming language**: Terraform test cases are written in [Go](https://golang.org/dl/).
- **dep**: [dep](https://github.com/golang/dep#installation) is a dependency management tool for Go.
- **Azure CLI**: The [Azure CLI](https://docs.microsoft.com/cli/azure/install-azure-cli?view=azure-cli-latest) is a command-line tool you can use to manage Azure resources. (Terraform supports authenticating to Azure through a service principal or [via the Azure CLI](https://www.terraform.io/docs/providers/azurerm/authenticating_via_azure_cli.html).)
- **mage**: We use the [mage executable](https://github.com/magefile/mage/releases) to show you how to simplify running Terratest cases. 

## Quickstart

```sh
dep ensure
az login
mage
```