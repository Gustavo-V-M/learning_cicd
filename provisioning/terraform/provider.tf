terraform {
  required_providers {
    huaweicloud = {
      source  = "huaweicloud/huaweicloud"
      version = ">= 1.36.0"
    }
  }
}

variable "secret_key" {}
variable "access_key" {}

provider "huaweicloud" {
    region = "sa-brazil-1"
    secret_key = var.secret_key
    access_key = var.access_key
}