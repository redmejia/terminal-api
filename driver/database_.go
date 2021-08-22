package driver

import "log"

func (d *dbRepo) Save() {
	err := d.db.Ping()
	if err != nil {
		log.Println("erroro ", err)
	}

	log.Println("From Ok Ok")
}
