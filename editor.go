// Package editor for quick display, editing and saving of data
package editor

import (
	"database/sql"
	"fmt"
	"github.com/skvdmt/f"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"
	"time"
)

// currentDir is full path to package
var currentDir string

// DBC database connection settings
var DBC Connections

// editorSessionIDCookieName cookie name for editor session id
var editorSessionIDCookieName = "EditorSessionID"

// init get full path to package directory and make handler to api to save data
func init() {
	// set current dir
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Println("No caller information")
	}
	currentDir = path.Dir(filename)

	// make database connections map
	DBC = make(Connections)

	// setup handlers
	http.HandleFunc("/editor/api/save", Save)
}

// GetJavaScript returns the JavaScript code of the element control
func GetJavaScript() string {
	data, err := os.ReadFile(fmt.Sprintf("%s/javascript/editor.min.js", currentDir))
	if err != nil {
		log.Println(err)
	}
	return string(data)
}

// GetCSS returns the CSS styling of the controls
func GetCSS() string {
	data, err := os.ReadFile(fmt.Sprintf("%s/css/editor.min.css", currentDir))
	if err != nil {
		log.Println(err)
	}
	return string(data)
}

// GetSessionID return editor session id from cookies
func GetSessionID(w http.ResponseWriter, r *http.Request, db *sql.DB) string {
	for _, cookie := range r.Cookies() {
		if cookie.Name == editorSessionIDCookieName {
			return cookie.Value
		}
	}
	return makeSession(w, db)
}

// makeSession creates a session and returns its id
func makeSession(w http.ResponseWriter, db *sql.DB) string {
	// creates a unique session ID
	var sessionID string
	for {
		hash := f.GetHash()
		var cnt int
		err := db.QueryRow("select count(*) from `editor` where `session_id` = ?", hash).Scan(&cnt)
		if err != nil {
			log.Println(err)
		}
		if cnt == 0 {
			sessionID = hash
			break
		}
	}
	// writes session ID in cookies
	http.SetCookie(w, &http.Cookie{
		Name:     editorSessionIDCookieName,
		Value:    sessionID,
		Path:     "/",
		Expires:  time.Now().AddDate(1, 0, 0),
		SameSite: http.SameSiteLaxMode,
	})
	return sessionID
}
