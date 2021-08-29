package models

import "time"

type Project struct {
	ProjectID          int64     `json:"project_id"`
	DevID              int64     `json:"dev_id"`
	ProjectName        string    `json:"project_name"`
	ProjectDescription string    `json:"project_description"`
	Created            time.Time `json:"created"`
	CreatedBy          string    `json:"created_by"`
	ProjectRepo        string    `json:"project_repo"`
	ProjectLive        string    `json:"project_live"`
}
