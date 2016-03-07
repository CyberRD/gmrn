package apis

import (
//"encoding/json"
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
	WorkInProgress *bool   `json:"work_in_progress"`
}

type Project struct {
	Id                float64 `json:"id"`
	PathWithNamespace string  `json:"path_with_namespace"`
}
