package apis

import (
	//fmt"
	log "github.com/Sirupsen/logrus"
)

func InitGitlabApi(url, token string) *GitLabApi {
	log.Infof("Init Gitlab Api with token : %s", token)
	gitlabpai := GitLabApi{}
	gitlabpai.token = token
	return &gitlabpai
}

type GitLabApi struct {
	gitlaburl string
	token     string
}

func (gitlab *GitLabApi) GetProjects() *[]Poroject {
	return nil
}
