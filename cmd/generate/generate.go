package main

import (
	"github.com/MichaelThessel/photo/model"
	"github.com/MichaelThessel/photo/service"
)

func main() {
	ag := service.AlbumGenerator{
		ExportPath: "./data",
	}

	fg := service.FolderGenerator{
		ExportPath: "./data",
	}

	tree := model.NewTree()
	tree.Parse()

	for _, folder := range tree.Folders {
		fg.GenerateHTML(folder)
		for _, album := range folder.Albums {
			ag.GenerateHTML(folder, album)
		}
	}

}
