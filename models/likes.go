package models

type Likes struct {
	ProjectID int64 `json:"project_id"`
	LikeCount int64 `json:"like_count"`
}
