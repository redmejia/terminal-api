package driver

import "github.com/redmejia/terminal/models"

// ITerminal interface methods
type IDatabaseRepo interface {
	InsertNewDev(user models.User) error
	InsertNewProject(project models.Project) error
}
