# Example CLI

This is an example CLI you can use to exercise the functionality in this Plex client library.

## Executing

To execute this program, run:

```
go run .
```

This will print a list of commands that can be executed.

### Authentication

Before interacting with your server, you will need to provision an access token. This can be done by one of two options:

* **Set token** *(recommended)*: if you already know your access token (Plex [has documentation on how to find this](https://support.plex.tv/articles/204059436-finding-an-authentication-token-x-plex-token/)), you can initialize your access token by executing:

```
go run . token <token>
```

* **Username/password**: if you do not have [2FA enabled](https://support.plex.tv/articles/two-factor-authentication/) (and you should), if you execute the following, this will retrieve an access token for you:

```
go run . signin <username> <password>
```

Following that, you will need to select a server:

```
go run . pick-server
```

After this, you will be able to use the list of commands used to interact with your server.