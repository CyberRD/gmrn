package notifier

import (
	log "github.com/Sirupsen/logrus"
	"github.com/eternnoir/gmrn/apis"
	"strconv"
	"time"
)

func InitGitLabNotifier(url, token string, projects []string, pollingInterval, notifyInterval time.Duration) *GitLabNotifier {
	gitlab := GitLabNotifier{}
	gitlab.Url = url
	gitlab.Token = token
	gitlab.Projects = projects
	gitlab.PollingInterval = pollingInterval
	gitlab.NotifyInterval = notifyInterval
	gitlab.NotifyRunners = []NotifyRunner{}
	gitlab.Api = apis.InitGitlabApi(url, token)
	gitlab.MRLastNotifyTime = make(map[string]time.Time)
	log.Infof("Init GitLabNotifier. Url:%s , Toke:%s, %d projects.", url, token, len(projects))
	return &gitlab
}

type GitLabNotifier struct {
	Url              string
	Token            string
	Projects         []string
	PollingInterval  time.Duration
	NotifyInterval   time.Duration
	Api              *apis.GitLabApi
	NotifyRunners    []NotifyRunner
	MRLastNotifyTime map[string]time.Time
}

func (notifier *GitLabNotifier) Run() {

	notifier.CheckProjects()
	//  loops forever to polling merge request.
	for {
		err := notifier.notifyForMergeRequest()
		if err != nil {
			log.Error(err)
		}
		time.Sleep(notifier.PollingInterval)
	}
}
func (notifier *GitLabNotifier) AppendNotifyRunner(runner NotifyRunner) {
	log.Infof("Append Runner %#v", runner)
	notifier.NotifyRunners = append(notifier.NotifyRunners, runner)
}

func (notifier *GitLabNotifier) notifyForMergeRequest() error {
	allMrs, err := notifier.GetAllProjectsMr()
	if err != nil {
		log.Error(err)
		return err
	}
	log.Infof("Get %d Merge Requests", len(allMrs))
	for _, mr := range allMrs {
		go notifier.triggerNotifyCommand(mr)
	}
	return nil
}

func (notifier *GitLabNotifier) triggerNotifyCommand(mr *apis.MergeRequest) {
	if mr.WorkInProgress {
		log.Debugf("%s Merge Reques is WorkInProgress. Do not need to notify.", mr.Title)
		return
	}
	uumrid := strconv.Itoa(int(mr.ProjectId)) + ":" + strconv.Itoa(int(mr.Id))
	if val, ok := notifier.MRLastNotifyTime[uumrid]; ok {
		if time.Now().Before(val.Add(notifier.NotifyInterval)) {
			// Do not need to run notify command.
			log.Debugf("Do not need to run notify command for %s", mr.Title)
			return
		}
	}
	log.Infof("Trigger command for %s", mr.Title)
	notifier.MRLastNotifyTime[uumrid] = time.Now()
	notifier.runNotifyCommand(mr)

}

func (notifier *GitLabNotifier) runNotifyCommand(mr *apis.MergeRequest) {
	mrerr := mr.GetProjectInfo(notifier.Api)
	if mrerr != nil {
		log.Errorf("Try to get merge request's project detial [FAIL]. %#v", mr)
		return
	}

	for _, nr := range notifier.NotifyRunners {
		log.Infof("Start Trigger NotifyRunner %#v", nr)
		err := nr.Trigger(mr)
		if err != nil {
			log.Errorf("Trigger NotifyRunner %#v [FAIL]", nr)
		}
		log.Infof("Tirgger NotifyRunner %#v  [Success].", nr)
	}
}

func (notifier *GitLabNotifier) GetAllProjectsMr() ([]*apis.MergeRequest, error) {
	var mrs []*apis.MergeRequest
	for _, projectId := range notifier.Projects {
		resultmrs, err := notifier.Api.GetMergeRequests(projectId, "opened")
		if err != nil {
			log.Error(err)
			return nil, err
		}
		mrs = append(mrs, resultmrs...)
	}
	return mrs, nil
}

// checkProjects is check notifier's projects is exist or not.
// If projects is empty, It will set all project to project list.
func (notifier *GitLabNotifier) CheckProjects() error {
	if len(notifier.Projects) < 1 {
		log.Infof("Notifier' project is empty. Load all projects from gitlab.")
		err := notifier.setAllProjectId()
		if err != nil {
			return err
		}
		log.Infof("%d projects loaded.", len(notifier.Projects))
	}
	return nil
}

func (notifier *GitLabNotifier) setAllProjectId() error {
	projects, err := notifier.Api.GetProjects()
	if err != nil {
		log.Error(err)
		return err
	}
	for _, project := range projects {
		notifier.Projects = append(notifier.Projects, project.PathWithNamespace)
	}
	return nil
}
