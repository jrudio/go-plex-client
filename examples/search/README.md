Getting Started

Initialize Go modules by running (if not done already)

`go mod init <your-repo-name>`

Import `go-plex-client` in your code:

```
import (
    plex "github.com/jrudio/go-plex-client"
)
```
add imported packages to your `go.mod` by running:
`go mod tidy`

Build your Go program
`go build -o plex-client`

Run program
`./plex-client <media-title> -url http://localhost:32400 -token abc123`