terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "3.5.0"
    }
  }
}
 
provider "google" {
  credentials = "/home/project/.sa_key"
  project = "cal-3032-160494a73309"
  region = "us-central1"
}
 
module "webserver" {
    source = "../../"
    name         = var.name
    machine_type = var.machine_type
    zone         = var.zone
}