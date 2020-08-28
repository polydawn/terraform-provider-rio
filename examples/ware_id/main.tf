terraform {
  required_providers {
    rio = {
      versions = ["0.1"]
      source = "github.com/polydawn/rio"
    }
  }
}


data "rio_ware_id" "build" {
  type = "tar"
  hash = "9hubDhroT7DNbyT9FVfL7iafpA9Apd1bK1REQ1Ldig2dSfMDvoBru4LhtXg9PQ7RNr"
}

output "build_ware_id" {
  value = data.rio_ware_id.build.id
}

output "build_ware_type" {
  value = data.rio_ware_id.build.type
}

output "build_ware_hash" {
  value = data.rio_ware_id.build.hash
}

data "rio_parse_ware_id" "parse" {
  id = data.rio_ware_id.build.id
}

output "parse_ware_id" {
  value = data.rio_parse_ware_id.parse.id
}

output "parse_ware_type" {
  value = data.rio_parse_ware_id.parse.type
}

output "parse_ware_hash" {
  value = data.rio_parse_ware_id.parse.hash
}
