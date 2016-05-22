# Plex.tv and Plex Media Server client written in Go

`go get -u github.com/jrudio/go-plex-client`

### Usage

```Go
Plex, err := plex.New("http://192.168.1.2:32400", "myPlexToken")

// Test your connection to your Plex server
result, pErr := Plex.Test()

// Search for media in your plex server
results, pErr := Plex.Search("The Walking Dead")

// ... and more! Please checkout plex.go for more methods
```

#### The following require a plex token

Get friends: `GET https://plex.tv/api/users`

Delete a friend: `DELETE https://plex.tv/api/friends/<friend-id>`

Get servers: `GET https://plex.tv/api/servers`

Get secure address to your server: `GET https://plex.tv/pms/resources.xml?includeHttps=1`

Use secure host to your server `https://<server-ip>.<hash>.plex.direct:<port>/`

Get server libraries / section ids: `GET https://plex.tv/api/servers/<machine-id>`

Get server libraries / section info: `GET <ip>:<port>/library/sections`

Validate Plex username or email is an email: `POST https://plex.tv/api/users/validate?invited_email=<username-or-email>`

Get labels: `GET <ip>:32400/library/sections/<section-id>/label?type=<library-index>`

Get media that is assigned to label: `GET /library/sections/<library-key>/all?label=<label-id>&type=<library-index>&sort=titleSort%3Aasc&X-Plex-Container-Start=0&X-Plex-Container-Size=20`

Invite friend: `POST https://plex.tv/api/servers/<machine-id>/shared_servers`

Example post body
```json
{
  "server_id": "abc123",
  "shared_server": {
    "library_section_ids": [11247160],
    "invited_email": "jrudio-guest"
  },
  "sharing_settings": {
    "filterMovies": "label=jrudio-guest"
  }
}
```

Get invites: `GET https://plex.tv/api/invites/requests`

Accept invite: `https://plex.tv/api/invites/requests/<request-id>?friend=0&server=1&home=0`

Get invites that were sent / pending: `GET https://plex.tv/api/invites/requested`