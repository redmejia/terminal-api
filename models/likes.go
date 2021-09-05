package models

// Likes
type Likes struct {
	ProjectID int64 `json:"project_id"`
	LikeCount int64 `json:"like_count"`
}

func (l *Likes) IsLiked() bool {
	// check if like count is 0
	if l.LikeCount > 0 {
		return true
	}

	return false
}
