package driver

import (
	"net/http"

	"github.com/redmejia/terminal/models"
)

// ITerminal interface methods
type IDatabaseRepo interface {
	GetProjects() ([]models.Project, error)
	GetProjectById(projectId int64) (models.Project, error)

	// InsertNewDev(user models.User) error
	InsertNewDev(user models.User, w http.ResponseWriter) error
	InsertNewProject(project models.Project) error

	DeleteProject(projectId, devId int64) error

	UpdateProject(p models.Project) error

	LikeAProject(projectId, devId int64) error

	MakeAComment(comment models.Comments) error
}
