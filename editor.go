// Package editor for quick display, editing and saving of data
package editor

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"
)

// currentDir is full path to package
var currentDir string

var DBC Connections

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
