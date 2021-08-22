package driver

import "log"

func (d *dbRepo) Save() {
	err := d.conn.Ping()
	if err != nil {
		log.Println("erroro ", err)
	}

	log.Println("From Ok Ok")
}
