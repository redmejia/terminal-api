package models

type Project struct {
	DevID       int64  `json:"dev_id"`
	ProjectName string `json:"project_name"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by"`
	Repo        string `json:"repo"`
	Live        string `json:"live"`
}
