package driver

import "github.com/redmejia/terminal/models"

// ITerminal interface methods
type IDatabaseRepo interface {
	GetProjects() ([]models.Project, error)
	GetProjectById(projectId int64) (models.Project, error)

	InsertNewDev(user models.User) error
	InsertNewProject(project models.Project) error

	DeleteProject(projectId, devId int64) error

	UpdateProject(p models.Project) error

	LikeAProject(projectId, devId int64) error
}
