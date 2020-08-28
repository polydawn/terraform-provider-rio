terraform {
  required_providers {
    rio = {
      versions = ["0.1"]
      source = "github.com/polydawn/rio"
    }
  }
}


provider "rio" {}

module "scan" {
  source = "./scan"
}

output "scan_id" {
  value = module.scan.id
}

output "scan_type" {
  value = module.scan.type
}

output "scan_path" {
  value = module.scan.path
}
