package editor

import (
	"database/sql"
	"log"
)

// Connections to databases
type Connections map[int]string

// Add connection
func (c *Connections) Add(id int, source string) {
	(*c)[id] = source
}

func (c *Connections) Open(id int) *sql.DB {
	db, err := sql.Open("mysql", (*c)[id])
	if err != nil {
		log.Println(err)
	}
	return db
}

func (c *Connections) Close(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Println(err)
	}
}
