package editor

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// SaveResponse is json data for ApiSave response
type SaveResponse struct {
	Changed bool   `json:"changed"`
	Error   string `json:"error,omitempty"`
}

// Save is API connect to database (need setup before) to save data by send key
func Save(w http.ResponseWriter, r *http.Request) {
	// read incoming data
	var data struct {
		ConnId int         `json:"conn-id"`
		Key    string      `json:"key"`
		Value  interface{} `json:"value"`
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		log.Println(err)
	}
	var res SaveResponse
	// if ConnID exists
	if _, ok := DBC[data.ConnId]; ok {
		// connect to db
		db := DBC.Open(data.ConnId)
		defer DBC.Close(db)
		// delete old data lines
		_, err = db.Exec(fmt.Sprintf("delete from `editor` where `dt_last_use` < date_sub(now(),interval 1 hour)"))
		if err != nil {
			log.Println(err)
		}
		// getting a table, column, and editable data field ID by key
		var table, row, dt string
		var idl int
		err = db.QueryRow("select `table`, `row_name`, `id_line`, `type` from `editor` where `key` = ?", data.Key).Scan(&table, &row, &idl, &dt)
		if err == nil {
			// update date time last use current key
			_, err = db.Exec("update `editor` set `dt_last_use` = CURRENT_TIMESTAMP where `key` = ?", data.Key)
			if err != nil {
				log.Println(err)
			}
			// updating data in accordance with the information received by the key
			switch dt {
			case "input-text", "content-html", "textarea", "select":
				_, err := db.Exec(fmt.Sprintf("update `%s` set `%s` = ? where `id` = ?", table, row), data.Value.(string), idl)
				if err != nil {
					log.Println(err)
				}
			case "checkbox":
				_, err = db.Exec(fmt.Sprintf("update `%s` set `%s` = ? where `id` = ?", table, row), data.Value.(bool), idl)
				if err != nil {
					log.Println(err)
				}
			}
			res.Changed = true
		} else if err.Error() == "sql: no rows in result set" {
			// key missing
			res.Error = "wrong_key"
		} else {
			// another error
			res.Error = err.Error()
			log.Fatalln(err)
		}
	} else {
		res.Error = "wrong_connID"
	}
	// json response formatting
	jr, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}
	// sending header and response to the operation in JSON format
	w.Header().Set("Content-type", "application/json")
	_, err = w.Write(jr)
	if err != nil {
		log.Println(err)
	}
}
