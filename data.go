package editor

import (
	"bytes"
	"fmt"
	"github.com/skvdmt/f"
	"log"
)

// NewData returns a reference to a new type with data
func NewData(connID int, table, row string, id int, datatype string, value interface{}) *Data {
	d := &Data{
		ConnID:   connID,
		table:    table,
		row:      row,
		id:       id,
		DataType: datatype,
		Value:    value,
		Editable: true,
		Key:      f.GetHash(),
	}
	return d
}

// Data data type for quick editing
type Data struct {
	ConnID   int
	table    string
	row      string
	id       int
	DataType string
	Value    interface{}
	Editable bool
	Key      string
}

// Edit preparation of a template that allows for quick editing on the page
func (d *Data) Edit() string {
	_, err := DBC[d.ConnID].Query("insert into `editor` (`table`, `row_name`, `id_line`, `type`, `key`) VALUES (?, ?, ?, ?, ?)", d.table, d.row, d.id, d.DataType, d.Key)
	if err != nil {
		log.Println(err)
	}
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
