package models

// Comments
type Comments struct {
	ProjectID   int64  `json:"project_id"`
	DevUsername string `json:"dev_username"`
	DevComment  string `json:"dev_comment"`
}
