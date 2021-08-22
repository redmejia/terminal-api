package driver

import "log"

func (d *DbRepo) Save() {
	err := d.Conn.Ping()
	if err != nil {
		log.Println("erroro ", err)
	}

	log.Println("From Ok Ok")
}
