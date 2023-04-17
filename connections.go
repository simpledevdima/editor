package editor

import (
	"database/sql"
	"fmt"
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
	fmt.Println("Open")
	if err != nil {
		log.Println(err)
	}
	return db
}

func (c *Connections) Close(db *sql.DB) {
	err := db.Close()
	fmt.Println("Close")
	if err != nil {
		log.Println(err)
	}
}
