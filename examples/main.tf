terraform {
  required_providers {
    rio = {
      versions = ["0.1"]
      source = "github.com/polydawn/rio"
    }
  }
}


provider "rio" {}

module "tar" {
  source = "./tar"
}

output "test" {
  value = module.tar.test
}