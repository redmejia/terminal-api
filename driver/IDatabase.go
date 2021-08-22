package driver

import "github.com/redmejia/terminal/models"

// ITerminal interface methods
type IDatabaseRepo interface {
	Save(user models.User) error
}
