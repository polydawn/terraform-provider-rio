terraform {
  required_providers {
    rio = {
      versions = ["0.1"]
      source = "github.com/polydawn/rio"
    }
  }
}

data "rio_pack" "pack" {
  type = "tar"
  path = abspath(path.root)
  target = "file://${abspath(path.root)}/output.tar"

  filters = {
    mtime = "2020-08-28T01:04:00-05:00"
    gid = 2000
  }
}

output "id" {
  value = data.rio_pack.pack.id
}

output "type" {
  value = data.rio_pack.pack.type
}

output "path" {
  value = data.rio_pack.pack.path
}
