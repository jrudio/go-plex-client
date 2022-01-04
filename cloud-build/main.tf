provider "google" {
  credentials = file("~/service-accounts/abc.json") # todo: use env variable
  project = "pms-environment" # todo: use env variable
  region = "us-west1" # this is cheaper than using Los Angeles
}

provider "google_compute_instance" "default" {
  name = "pms-linux-latest"
  machine_type = "e2-micro"
  zone = "us-west1-a"

  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-1604-xenial-v20170202"
    }
  }
}