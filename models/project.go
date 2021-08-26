package models

type Project struct {
	DevID              int64  `json:"dev_id"`
	ProjectName        string `json:"project_name"`
	ProjectDescription string `json:"project_description"`
	CreatedBy          string `json:"created_by"`
	Repo               string `json:"project_repo"`
	Live               string `json:"project_live"`
}
