package notifier

import (
	log "github.com/Sirupsen/logrus"
	"github.com/eternnoir/gmrn/apis"
	"time"
)

func InitGitLabNotifier(url, token string, projects []string, interval time.Duration) *GitLabNotifier {
	gitlab := GitLabNotifier{}
	gitlab.Url = url
	gitlab.Token = token
	gitlab.Projects = projects
	gitlab.Interval = interval
	gitlab.Api = apis.InitGitlabApi(url, token)
	log.Infof("Init GitLabNotifier. Url:%s , Toke:%s, %d projects.", url, token, len(projects))
	return &gitlab
}

type GitLabNotifier struct {
	Url      string
	Token    string
	Projects []string
	Interval time.Duration
	Api      *apis.GitLabApi
}
