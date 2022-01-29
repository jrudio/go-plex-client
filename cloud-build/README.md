CI/CD Flow
==

1. run go test

1. build go-plex-client

1. push to Artifact Registry

1. test cli against a docker instance of Plex Media Server

  - startup script should download test mp4 from cloud storage
  - link throw away Plex account to this server


test go-plex-client against PMS server

test go-plex-client against Plex.tv api

### Reporting

Use PubSub to listen for build status and report via Cloud Function to either:

- badge in repository
- selected PR

Resources:

[1] https://cloud.google.com/build/docs/building/build-go#configuring_go_builds

[2] https://cloud.google.com/build/docs/configuring-builds/pass-data-between-steps

[3] https://cloud.google.com/build/docs/subscribe-build-notifications