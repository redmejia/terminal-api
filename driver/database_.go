package driver

import (
	"time"

	"github.com/redmejia/terminal/models"
)

// InsertNewDev insert new developer to database
func (d *dbRepo) InsertNewDev(user models.User) error {
	var err error
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	row := tx.QueryRow(`INSERT INTO register (dev_email) VALUES($1) RETURNING dev_id`, user.DevEmail)

	var devId int64
	err = row.Scan(&devId)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`INSERT INTO signin (dev_id, dev_email, dev_password) VALUES ($1, $2, $3)`,
		devId, user.DevEmail, user.DevPassword,
	)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return err
}

// Insert new developer project
func (d *dbRepo) InsertNewProject(project models.Project) error {

	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	row := tx.QueryRow(`
		INSERT INTO projects (dev_id, created, created_by, project_name, project_description)
		VALUES ($1, $2, $3, $4, $5) RETURNING project_id
	`, project.DevID, time.Now(), project.CreatedBy, project.ProjectName, project.ProjectDescription,
	)

	var projectId int64
	err = row.Scan(&projectId)
	if err != nil {
		return err
	}

	// if reposotory or live project links
	_, err = tx.Exec(`
		INSERT INTO links (project_id, dev_id, project_repo, project_live)
		VALUES ($1, $2, $3, $4)
	`, projectId, project.DevID, project.ProjectRepo, project.ProjectLive,
	)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// DeleteProject
func (d *dbRepo) DeleteProject(projectId, devId int64) error {

	_, err := d.db.Exec(`DELETE FROM projects WHERE project_id = $1 AND dev_id = $2`, projectId, devId)
	if err != nil {
		return err
	}

	return nil

}

// UpdateProject
func (d *dbRepo) UpdateProject(project models.Project) error {

	tx, err := d.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.Exec(`
		UPDATE projects
			SET project_name = $3, project_description = $4
		WHERE project_id = $1 AND dev_id = $2`,
		project.ProjectID, project.DevID, project.ProjectName, project.ProjectDescription)
	if err != nil {
		return err
	}

	_, err = d.db.Exec(`
		UPDATE links
			SET project_repo = $3, project_live = $4
		WHERE project_id = $1 AND dev_id = $2`,
		project.ProjectID, project.DevID, project.ProjectRepo, project.ProjectLive,
	)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
