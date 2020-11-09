terraform {
  backend "remote" {
    hostname = "app.terraform.io"
    organization = "devops-mindset"

    workspaces {
      name = "Concord"
    }
  }
}

# Configure the Azure Provider
provider "azurerm" {
  # whilst the `version` attribute is optional, we recommend pinning to a given version of the Provider
  version = "=2.34.0"
  features {}
}

module "full-webapp" {
  source  = "app.terraform.io/devops-mindset/full-webapp/azurerm"
  version = "0.0.5"
  name = var.name
  region = var.region
}

