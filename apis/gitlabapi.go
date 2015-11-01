package apis

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/eternnoir/gmrn/utils"
	"net/url"
)

func InitGitlabApi(url, token string) *GitLabApi {
	log.Infof("Init Gitlab Api with token : %s", token)
	gitlabpai := GitLabApi{}
	gitlabpai.gitlaburl = url
	gitlabpai.token = token
	gitlabpai.apiversion = "v3" // Temporarily use v3 as default.
	return &gitlabpai
}

type GitLabApi struct {
	gitlaburl  string
	token      string
	apiversion string
}

func (gitlab *GitLabApi) GetProjects() ([]*Project, error) {
	gitlaburl := gitlab.gitlaburl
	token := gitlab.token
	log.Infof("Start get project list from %s use token %s", gitlaburl, token)
	params := url.Values{}
	params.Add("private_token", token)
	requestUrl := gitlab.GenApiUrl("projects")
	log.Infof("Send Request:%s", requestUrl)
	result, err := utils.SendGetRequest(requestUrl, params)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	resultStr := string(result[:])
	log.Debugf("Get Result %s", resultStr)

	// DeJosn
	var projects []*Project
	err = json.Unmarshal(result, &projects)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Infof("Get %d projects.", len(projects))
	return projects, nil
}

func (gitlab *GitLabApi) GenApiUrl(method string) string {
	url := fmt.Sprintf("%s/api/%s/%s", gitlab.gitlaburl, gitlab.apiversion, method)
	return url
}
