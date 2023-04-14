package editor

import "database/sql"

// Connections to databases
type Connections map[int]*sql.DB

// Add connection
func (c *Connections) Add(id int, conn *sql.DB) {
	(*c)[id] = conn
}
