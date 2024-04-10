package service

import (
	"html/template"
	"log"
	"os"

	"github.com/MichaelThessel/photo/model"
)

type FolderGenerator struct {
	ExportPath string
}

func (ag *FolderGenerator) GenerateHTML(folder model.Folder) {
	folderPath := ag.ExportPath + "/" + folder.Name

	f, err := os.Create(folderPath + "/index.html")
	if err != nil {
		log.Println("Coudn't open folder index file", err)
	}
	defer f.Close()

	tmpl, err := template.ParseFiles(
		"templates/folder.html",
		"templates/head.html",
		"templates/styles.html",
		"templates/header.html",
		"templates/footer.html",
	)
	if err != nil {
		log.Panic("Coudn't open template files", err)
	}

	templateData := struct {
		Page   model.Page
		Folder model.Folder
	}{
		Page: model.Page{
			Title: folder.Name,
		},
		Folder: folder,
	}
	err = tmpl.Execute(f, templateData)
	if err != nil {
		log.Panic("Coudn't compile template", err)
	}
}
