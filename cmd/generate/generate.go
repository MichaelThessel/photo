package main

import (
	"github.com/MichaelThessel/photo/model"
	"github.com/MichaelThessel/photo/service"
)

func main() {
	ag := service.AlbumGenerator{
		ExportPath: "./data",
	}

	tree := model.NewTree()
	tree.Parse()

	for _, folder := range tree.Folders {
		for _, album := range folder.Albums {
			ag.GenerateAlbumHTML(folder, album)
		}
	}

}
