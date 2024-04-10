package main

import (
	"html/template"
	"log"
	"os"

	"github.com/MichaelThessel/photo/model"
)

func main() {
	ag := AlbumGenerator{
		exportPath: "./data",
	}

	tree := model.NewTree()
	tree.Parse()

	for _, folder := range tree.Folders {
		for _, album := range folder.Albums {
			ag.GenerateAlbumHTML(folder, album)
		}
	}

}

type AlbumGenerator struct {
	exportPath string
}

func (ag *AlbumGenerator) GenerateAlbumHTML(folder model.Folder, album model.Album) {
	albumPath := ag.exportPath + "/" + folder.Name + "/" + album.Name

	f, err := os.Create(albumPath + "/index.html")
	if err != nil {
		log.Println("Coudn't open album index file", err)
	}
	defer f.Close()

	tmpl, err := template.ParseFiles(
		"templates/album.html",
		"templates/head.html",
		"templates/header.html",
		"templates/footer.html",
	)
	if err != nil {
		log.Panic("Coudn't open template files", err)
	}

	templateData := struct {
		Page  model.Page
		Album model.Album
	}{
		Page: model.Page{
			Title: album.Name,
		},
		Album: album,
	}
	err = tmpl.Execute(f, templateData)
	if err != nil {
		log.Panic("Coudn't compile template", err)
	}
}
