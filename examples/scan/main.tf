terraform {
  required_providers {
    rio = {
      versions = ["0.1"]
      source = "github.com/polydawn/rio"
    }
  }
}


data "rio_pack" "scan" {
  type = "tar"
  path = abspath(path.root)
}

output "id" {
  value = data.rio_pack.scan.id
}

output "type" {
  value = data.rio_pack.scan.type
}

output "path" {
  value = data.rio_pack.scan.path
}
