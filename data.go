package editor

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/skvdmt/f"
	"log"
)

// NewData returns a reference to a new type with data
func NewData(connID int, sessionID, table, row string, id int, datatype string, value interface{}) *Data {
	d := &Data{
		ConnID:    connID,
		table:     table,
		row:       row,
		id:        id,
		DataType:  datatype,
		Value:     value,
		Editable:  true,
		sessionID: sessionID,
	}
	return d
}

// Data data type for quick editing
type Data struct {
	sessionID string
	ConnID    int
	table     string
	row       string
	id        int
	DataType  string
	Value     interface{}
	Editable  bool
	Key       string
}

// Edit preparation of a template that allows for quick editing on the page
func (d *Data) Edit(db *sql.DB) string {
	d.getKey(db)
	var t bytes.Buffer
	tfn := fmt.Sprintf("%s/templates/%s.gohtml", currentDir, d.DataType)
	switch d.DataType {
	case "content-html":
		t = TemplateAsText(tfn, d)
	case "checkbox", "input-text", "textarea", "select":
		t = TemplateAsHTML(tfn, d)
	}
	return t.String()
}

// getKey checks if the current user record key exists in the database and returns it.
// Or creates and saves a new key and returns it
func (d *Data) getKey(db *sql.DB) {
	var sc int
	err := db.QueryRow("select count(*) from `editor` where `session_id` = ? and `table` = ? and `row_name` = ? and `id_line` = ? and `type` = ?", d.sessionID, d.table, d.row, d.id, d.DataType).Scan(&sc)
	if err != nil {
		log.Println(err)
	}
	switch sc {
	case 1:
		// get key
		err = db.QueryRow("select `key` from `editor` where `session_id` = ? and `table` = ? and `row_name` = ? and `id_line` = ? and `type` = ?", d.sessionID, d.table, d.row, d.id, d.DataType).Scan(&d.Key)
		if err != nil {
			log.Println(err)
		}
		// update last use
		_, err = db.Exec("update `editor` set `dt_last_use` = CURRENT_TIMESTAMP where `key` = ?", d.Key)
		if err != nil {
			log.Println(err)
		}
	case 0:
		// make new unique key
		for {
			d.Key = f.GetHash()
			var kc int
			err = db.QueryRow("select count(*) from `editor` where `key` = ?", d.Key).Scan(&kc)
			if err != nil {
				log.Println(err)
			}
			if kc == 0 {
				break
			}
		}
		// save key
		_, err = db.Exec("insert into `editor` (`session_id`, `table`, `row_name`, `id_line`, `type`, `key`) VALUES (?, ?, ?, ?, ?, ?)", d.sessionID, d.table, d.row, d.id, d.DataType, d.Key)
		if err != nil {
			log.Println(err)
		}
	}
}
