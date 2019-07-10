provider "aws" {
  shared_credentials_file = "$HOME/.aws/credentials"
  profile                 = "default"
  region                  = var.aws_region
}

data "aws_caller_identity" "current" {}

terraform {
  backend "s3" {
    bucket="rolli3net"
    key="terraform/terraform.tfstate"
    region="us-west-2"
  }
}
