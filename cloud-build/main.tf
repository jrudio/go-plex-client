provider "google" {
  region = "us-west1" # this is cheaper than using Los Angeles
}

resource "google_storage_bucket" "media-bucket" {
  name = "pms-media"
  location = "us-west1"
  storage_class = "STANDARD"

  versioning {
    enabled = true
  }

  labels = {
    "env" = "dev"
    "go-plex-client-bucket" = "guest"
  }
}