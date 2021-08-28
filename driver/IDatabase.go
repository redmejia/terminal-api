package driver

import "github.com/redmejia/terminal/models"

// ITerminal interface methods
type IDatabaseRepo interface {
	InsertNewDev(user models.User) error
	InsertNewProject(project models.Project) error
	DeleteProject(projectId, devId int64) error
	UpdateProject(p models.Project) error
}
