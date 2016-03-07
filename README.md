# GMRN

The polling base gitlab merge request notifier.

## Quick Start

### Installation

* Use go get to install.

```
$ go get github.com/eternnoir/gmrn
```

* Download binary file from [Release Page](https://github.com/eternnoir/gmrn/releases).

### Create config file.

Config file use [toml](https://github.com/toml-lang/toml) format.

```toml
Url = "http://your.gitlab.url"
Token = "your_secret_token"
PollingInterval = "5s"
NotifyInterval= "15m30s"
Projects= ["eternnoir/gmrn","balabala/coolproject"] # Optional. Leave empty to monitor all projects.

[[CommandNotifyRunner]]
NotifyCommand = "/tmp/run.sh"

[[MMNotifyRunner]]
WebhookUrl = ""
Channel = ""
Username = "地方的Merge Rquest"
TextTemplate = """ {{if .Assignee}} @{{.Assignee.UserName}} {{else}} @channel {{end}} 地方的Merge Request 需要你的協助
* Project : {{.Project.PathWithNamespace}}
* Titel : {{.Title}}
* Author : {{.Author.UserName}}
* [{{.Project.WebUrl}}/merge_requests/{{.Iid}}]({{.Project.WebUrl}}/merge_requests/{{.Iid}})

請盡快改成 `WIP:` 不然我會一直呼叫你 ╮(╯◇╰)╭
"""

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

### Run

```bash
$ gmrn -c config.toml
```

You also can use -d flag to show debug level's log.

```bash
$ gmrn -d -c config.toml
```

## What kind of merge request will be notify?

* Opened
* Not in WIP(Work In Progress)
