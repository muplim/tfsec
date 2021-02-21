terraform {
  required_version = "0.14.6"
}

module "first-level" {
  source = "../../first-level"
}
