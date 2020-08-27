terraform {
  required_providers {
    rio = {
      versions = ["0.1"]
      source = "github.com/polydawn/rio"
    }
  }
}


data "rio_test" "good" {}

output "test" {
  value = data.rio_test.good.test
}


# data "rio_error" "bad" {}

# output "err" {
#   value = data.rio_error.bad.test
# }
