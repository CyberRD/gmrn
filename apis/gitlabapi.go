package apis

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/eternnoir/gmrn/utils"
	"net/url"
	"strings"
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

// GetProjects to get all projects on gitlab server.
func (gitlab *GitLabApi) GetProjects() ([]*Project, error) {
	gitlaburl := gitlab.gitlaburl
	token := gitlab.token
	log.Infof("Start get project list from %s use token %s", gitlaburl, token)
	params := url.Values{}
	params.Add("private_token", token)
	params.Add("page", "1")
	params.Add("per_page", "2000")
	requestUrl := gitlab.GenApiUrl("projects/all")
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

// GetProject can get specific project by project id or PathWithNamespace.
//
// id is like "1". PathWithNamesapce is lile "root/test2"
func (gitlab *GitLabApi) GetProject(id string) (*Project, error) {
	gitlaburl := gitlab.gitlaburl
	token := gitlab.token
	log.Infof("Start get project from %s use token %s", gitlaburl, token)
	params := url.Values{}
	params.Add("private_token", token)
	projectId := gitlab.GetRequestProjectId(id)
	requestUrl := gitlab.GenApiUrl("projects/" + projectId)
	log.Infof("Send Request:%s", requestUrl)
	result, err := utils.SendGetRequest(requestUrl, params)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	resultStr := string(result[:])
	log.Debugf("Get Result %s", resultStr)

	// DeJosn
	var project *Project
	err = json.Unmarshal(result, &project)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Infof("Get project %s", project.PathWithNamespace)
	return project, nil
}

// GetRequestProjectId is help to convert project id to request format.
// If projectId contain "/" must replace to "%2F"
func (gitlab *GitLabApi) GetRequestProjectId(projectId string) string {
	projectId = strings.Replace(projectId, "/", "%2F", -1)
	projectId = strings.Replace(projectId, ".", "%2E", -1)
	return projectId
}

func (gitlab *GitLabApi) GetMergeRequests(id, state string) ([]*MergeRequest, error) {
	token := gitlab.token
	log.Infof("Start get Merge Requests  from %s ", id)
	params := url.Values{}
	params.Add("private_token", token)
	params.Add("state", state)
	projectId := gitlab.GetRequestProjectId(id)
	requestUrl := gitlab.GenApiUrl(fmt.Sprintf("projects/%s/merge_requests", projectId))
	log.Infof("Send Request:%s", requestUrl)
	result, err := utils.SendGetRequest(requestUrl, params)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	resultStr := string(result[:])
	log.Debugf("Get Result %s", resultStr)
	var mergeRequests []*MergeRequest
	err = json.Unmarshal(result, &mergeRequests)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Infof("Get %d Merge Request from %s with %s state.", len(mergeRequests), id, state)
	return mergeRequests, nil
}

func (gitlab *GitLabApi) GenApiUrl(method string) string {
	url := fmt.Sprintf("%s/api/%s/%s", gitlab.gitlaburl, gitlab.apiversion, method)
	return url
}
