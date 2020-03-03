# DiscordLogin

An automated Discord token grabber, for use with [gtkcord3](https://github.com/diamondburned/gtkcord3).


## Why Webkit?

Using an actual browser allows us to solve Captchas normally, without needing JavaScript
hacks. Requests are done legitimately using the official website.

## How?

### Installation

```sh
# This should install to $GOPATH/bin/, or ~/go/bin/
go get github.com/diamondburned/discordlogin

# Add this to shellrc
PATH="$GOPATH/bin/:$PATH"
```

### Usage

```sh
TOKEN=$(discordlogin)
# Complete the rest in the browser. The subshell should finish when it has found
# the token.
```
