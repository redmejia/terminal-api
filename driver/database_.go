package driver

import (
	"time"

	"github.com/redmejia/terminal/models"
)

// GetProjectById retrive project by id
func (d *dbRepo) GetProjectById(projectId int64) (models.Project, error) {

	row := d.db.QueryRow(`
		SELECT p.dev_id,
			TO_CHAR(p.created, 'mon dy YYYY' ) AS created,
			p.created_by,
			p.project_name,
			p.project_description,
			l.project_repo,
			l.project_live
		FROM projects p
			JOIN links l ON p.project_id = l.project_id
		WHERE p.project_id = $1
	`, projectId)
	var project models.Project

	err := row.Scan(
		&project.DevID, &project.Created, &project.CreatedBy, &project.ProjectName,
		&project.ProjectDescription, &project.ProjectRepo, &project.ProjectLive,
	)
	if err != nil {
		return project, err // return empty project struct
	}

	return project, nil
}

// GetProjects retrive all projects
func (d *dbRepo) GetProjects() ([]models.Project, error) {
	var projects []models.Project

	rows, err := d.db.Query(`
		SELECT p.project_id,
			p.dev_id,
			TO_CHAR(p.created, 'mon dy YYYY' ) AS created,
			p.created_by,
			p.project_name,
			p.project_description,
			l.project_repo,
			l.project_live,
			ls.project_id,
			ls.like_count
		FROM projects p
			JOIN links l ON p.project_id = l.project_id
			LEFT JOIN likes ls ON p.project_id = ls.project_id
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Project
		rows.Scan(
			&p.ProjectID, &p.DevID, &p.Created, &p.CreatedBy, &p.ProjectName,
			&p.ProjectDescription, &p.ProjectRepo, &p.ProjectLive,
			&p.ProjectLike.ProjectID, &p.ProjectLike.LikeCount,
		)

		projects = append(projects, p)
	}

	return projects, nil

}

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

	// created := time.Unix(project.Created, 0)

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

// LikeAProject like a project 0 init like 1
func (d *dbRepo) LikeAProject(projectId, devId int64) error {
	const (
		like int64 = 1
		// topProjects       = 5 // set number of like to the top projects will be store this project on TOP PROJECT
	)

	row := d.db.QueryRow(`SELECT project_id, like_count FROM likes WHERE project_id = $1`, projectId)

	var likes models.Likes

	_ = row.Scan(&likes.ProjectID, &likes.LikeCount)

	if likes.IsLiked() {
		// update likes
		tx, err := d.db.Begin()
		if err != nil {
			return err
		}

		var newLike = likes.LikeCount + like

		_, err = tx.Exec(`
			UPDATE likes 
			SET like_count = $2 
			WHERE project_id = $1
		`, projectId, newLike)
		if err != nil {
			return err
		}

		_, err = tx.Exec(`
			INSERT INTO liked (project_id, dev_id, project_liked) VALUES ($1, $2, $3)
		`, projectId, devId, true)

		err = tx.Commit()
		if err != nil {
			return err
		}

	} else {
		// Initial likes
		tx, err := d.db.Begin()
		if err != nil {
			return err
		}

		_, err = tx.Exec(`
			INSERT INTO likes (project_id, like_count) VALUES ($1, $2)
		`, projectId, like)
		if err != nil {
			return err
		}

		_, err = tx.Exec(`
			INSERT INTO liked (project_id, dev_id, project_liked) VALUES ($1, $2, $3)
		`, projectId, devId, true)

		err = tx.Commit()
		if err != nil {
			return err
		}

	}
	return nil
}

// MakeAComment
func (d *dbRepo) MakeAComment(comment models.Comments) error {

	_, err := d.db.Exec(`
		 INSERT INTO comments (project_id, dev_username, dev_comment)
		 VALUES ($1, $2, $3)
	`, comment.ProjectID, comment.DevUsername, comment.DevComment,
	)
	if err != nil {
		return err
	}

	return nil
}
