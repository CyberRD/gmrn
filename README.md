# GMRN

The polling base gitlab merge request notifier. With Go/Golang.

## Quick Start

### Installation

Use go get to install.

```
$ go get github.com/eternnoir/gmrn
```

### Create config file.

Config file use (toml)[https://github.com/toml-lang/toml] format.

```toml
Url = "http://your.gitlab.url"
Token = "your_secret_token"
PollingInterval = "5s"
NotifyInterval= "30s"
NotifyCommand = "/tmp/run.sh"

```

### Run

```bash
$ gmrn -c config.toml
```

You also can use -d flag to show debug level's log.

```bash
$ gmrn -d -c config.toml
```
