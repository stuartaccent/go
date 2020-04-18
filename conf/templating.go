package conf

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ParseTemplates(root string) *template.Template {
	templ := template.New("")
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			_, err = templ.ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}
		return err
	})

	if err != nil {
		panic(err)
	}

	return templ
}
