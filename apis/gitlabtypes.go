package apis

import (
	//"encoding/json"a
	"strconv"
)

type Author struct {
	Name     string  `json:"name"`
	UserName string  `json:"username"`
	Id       float64 `json:"id"`
}
type MRList struct {
}

type MergeRequest struct {
	Id             float64 `json:"id"`
	Iid            float64 `json:"iid"`
	ProjectId      float64 `json:"project_id"`
	Title          string  `json:"title"`
	Author         *Author `json:"author"`
	State          string  `json:"state"`
	Assignee       *Author `json:"assignee"`
	WorkInProgress bool    `json:"work_in_progress"`
	Project        *Project
}

func (mr *MergeRequest) GetProjectInfo(api *GitLabApi) error {
	project, err := api.GetProject(strconv.Itoa(int(mr.ProjectId)))
	if err != nil {
		return err
	}
	mr.Project = project
	return nil
}

type Project struct {
	Id                float64 `json:"id"`
	PathWithNamespace string  `json:"path_with_namespace"`
	WebUrl            string  `json:"web_url"`
}
