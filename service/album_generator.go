package service

import (
	"html/template"
	"log"
	"os"

	"github.com/MichaelThessel/photo/model"
)

type AlbumGenerator struct {
	ExportPath string
}

func (ag *AlbumGenerator) GenerateHTML(folder model.Folder, album model.Album) {
	albumPath := ag.ExportPath + "/" + folder.Name + "/" + album.Name

	f, err := os.Create(albumPath + "/index.html")
	if err != nil {
		log.Println("Coudn't open album index file", err)
	}
	defer f.Close()

	tmpl, err := template.ParseFiles(
		"templates/album.html",
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
		Album  model.Album
		Folder model.Folder
	}{
		Page: model.Page{
			Title: album.Name,
		},
		Album:  album,
		Folder: folder,
	}
	err = tmpl.Execute(f, templateData)
	if err != nil {
		log.Panic("Coudn't compile template", err)
	}
}
