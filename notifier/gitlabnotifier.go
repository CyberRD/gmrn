package notifier

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/eternnoir/gmrn/apis"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func InitGitLabNotifier(url, token string, projects []string, pollingInterval, notifyInterval time.Duration, notifyCommand string) *GitLabNotifier {
	gitlab := GitLabNotifier{}
	gitlab.Url = url
	gitlab.Token = token
	gitlab.Projects = projects
	gitlab.PollingInterval = pollingInterval
	gitlab.NotifyInterval = notifyInterval
	gitlab.NotifyCommand = notifyCommand
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
	NotifyCommand    string
	MRLastNotifyTime map[string]time.Time
}

func (notifier *GitLabNotifier) Run() {

	notifier.checkProjects()
	//  loops forever to polling merge request.
	for {
		err := notifier.notifyForMergeRequest()
		if err != nil {
			log.Error(err)
		}
		time.Sleep(notifier.PollingInterval)
	}
}

func (notifier *GitLabNotifier) notifyForMergeRequest() error {
	allMrs, err := notifier.getAllProjectsMr()
	if err != nil {
		log.Error(err)
		return err
	}
	log.Infof("Get %d Merge Requests", len(allMrs))
	for _, mr := range allMrs {
		go notifier.triggerNitifyCommand(mr)
	}
	return nil
}

func (notifier *GitLabNotifier) triggerNitifyCommand(mr *apis.MergeRequest) {
	if mr.WorkInProgress != nil {
		log.Debugf("%s Merger Reques is WorkInProgress. Do not need to notify.", mr.Title)
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
	log.Infof("Start Trigger Command %s", notifier.NotifyCommand)
	binary, err := exec.LookPath(notifier.NotifyCommand)
	if err != nil {
		log.Error(err)
		return
	}
	cmd := exec.Command(binary)
	env := os.Environ()
	env = append(env, fmt.Sprintf("MR_TITLE=%s", mr.Title))
	cmd.Env = env
	cmdOut, _ := cmd.StdoutPipe()

	startErr := cmd.Start()
	if startErr != nil {
		log.Error(startErr)
		return
	}

	// read stdout and stderr
	stdOutput, _ := ioutil.ReadAll(cmdOut)
	log.Debug(string(stdOutput[:]))

	err = cmd.Wait()
	if err != nil {
		log.Error(err)
		return
	}
	log.Infof("Command %s Done.", notifier.NotifyCommand)
}

func (notifier *GitLabNotifier) getAllProjectsMr() ([]*apis.MergeRequest, error) {
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
func (notifier *GitLabNotifier) checkProjects() error {
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
