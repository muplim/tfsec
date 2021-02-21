terraform {
  required_version = "0.14.6"
}

module "second-level" {
  source = "../second-level"
}
