package controllers

import (
	"html/template"
	"github.com/vsmejkal/events/config"
)


func compileTemplate(layout, content string) (*template.Template, error) {
	path := config.Admin.DocumentRoot + "/templates/"

	tpl := template.New(content)
	tpl, err := tpl.ParseFiles(
		path + "layouts/" + layout + ".tmpl",
		path + content + ".tmpl",
		path + "pagination.tmpl",
	)
	if err != nil {
		return nil, err
	}
	return tpl, nil
}

