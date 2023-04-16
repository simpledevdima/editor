package editor

import (
	"bytes"
	"html/template"
	"log"
)

func TemplateAsHTML(file string, d *Data) bytes.Buffer {
	tpl, err := template.ParseFiles(file)
	if err != nil {
		log.Println(err)
	}
	var t bytes.Buffer
	err = tpl.Execute(&t, d)
	if err != nil {
		log.Println(err)
	}
	return t
}
