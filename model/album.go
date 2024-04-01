package model

import (
	"log"
	"os"
)

type Album struct {
	Name      string
	ThumbPath string
	Slides    []Slide
}

// CreateAlbumDirectories creates directories for an album to store images and
// thumbs
func (a Album) CreateAlbumDirectories(basePath string) {
	// Create album directories
	err := os.Mkdir(basePath+"/"+a.Name, 0744)
	if err != nil {
		log.Println("Error creating album directory", err)
	}

	err = os.Mkdir(basePath+"/"+a.Name+"/slides", 0744)
	if err != nil {
		log.Println("Error creating slide directory", err)
	}

	err = os.Mkdir(basePath+"/"+a.Name+"/thumbs", 0744)
	if err != nil {
		log.Println("Error creating thumb directory", err)
	}
}
