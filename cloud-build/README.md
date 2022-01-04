CI/CD Flow
==

1. run go test

1. build go-plex-client

1. push to Artifact Registry

1. test cli against a docker instance of Plex Media Server


boot up PMS-linux-container
- startup script should download test mp4 from cloud storage
- link throw away Plex account to this server

test go-plex-client against PMS server

test go-plex-client against Plex.tv api

Resources:

https://cloud.google.com/build/docs/building/build-go#configuring_go_builds
https://cloud.google.com/build/docs/configuring-builds/pass-data-between-steps