# GMRN

The polling base gitlab merge request notifier. With Go/Golang.

## Quick Start

### Installation

Use go get to install.

```
$ go get github.com/eternnoir/gmrn
```

### Create config file.

Config file use [toml](https://github.com/toml-lang/toml) format.

```toml
Url = "http://your.gitlab.url"
Token = "your_secret_token"
PollingInterval = "5s"
NotifyInterval= "15m30s"
Projects= ["eternnoir/gmrn","balabala/coolproject"] # Optional. Leave empty to monitor all projects.
NotifyCommand = "/tmp/run.sh"
```

* Url : Your gitlab site.
* Token : Secret token for your gitlab site. 
* PollingInterval : Interval between each polling request.
* NotifyInterval : Interval trigger NotifyCommand for each merge request.
* Projects : Merge Request in these projects, gmrn will notify.
* NotifyCommand : When get merge requests will call this command.

#### How to get token?

```
curl http://your.gitlab.site/api/v3/session --data 'login=myUser&password=myPass'
```
### What kind of merge request will be notify?

* Open
* Not in WIP(Work In Progress)

### Run

```bash
$ gmrn -c config.toml
```

You also can use -d flag to show debug level's log.

```bash
$ gmrn -d -c config.toml
```
