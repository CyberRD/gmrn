package notifier

import (
	//log "github.com/Sirupsen/logrus"
	"github.com/eternnoir/gmrn/apis"
	"time"
)

func InitGitLabNotifier(url, token string, interval time.Duration) *GitLabNotifier {
	gitlab := GitLabNotifier{}
	gitlab.Url = url
	gitlab.Token = token
	gitlab.Interval = interval
	gitlab.Api = apis.InitGitlabApi(url, token)
	return &gitlab
}

type GitLabNotifier struct {
	Url      string
	Token    string
	Interval time.Duration
	Api      *apis.GitLabApi
}
