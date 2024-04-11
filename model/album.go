package model

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
)

type Album struct {
	Name       string
	ThumbIndex int
	Slides     []Slide
}

// CreateAlbumDirectories creates directories for an album to store images and
// thumbs
func (a Album) CreateAlbumDirectories(basePath string) {
	// Create album directories
	albumPath := basePath + "/" + a.Name
	if _, err := os.Stat(albumPath); err != nil {
		err := os.Mkdir(albumPath, 0744)
		if err != nil {
			log.Fatalln("Error creating album directory", err)
		}
	}

	slidePath := basePath + "/" + a.Name + "/slides"
	if _, err := os.Stat(slidePath); err != nil {
		err := os.Mkdir(slidePath, 0744)
		if err != nil {
			log.Fatalln("Error creating slide directory", err)
		}
	}

	thumbPath := basePath + "/" + a.Name + "/thumbs"
	if _, err := os.Stat(thumbPath); err != nil {
		err := os.Mkdir(thumbPath, 0744)
		if err != nil {
			log.Fatalln("Error creating thumb directory", err)
		}
	}
}

// Creates the data.json file containing the album information
func (a Album) CreateDataJSON(basePath string) {
	json, err := json.MarshalIndent(a, "", "    ")

	if err != nil {
		log.Println("Error creating album JSON", err)
	}

	w, err := os.Create(basePath + "/" + a.Name + "/data.json")
	if err != nil {
		log.Println("Couldn't create data.json", err)
		return
	}
	defer w.Close()

	_, err = w.ReadFrom(bytes.NewReader(json))
	if err != nil {
		log.Println("Couldn't write data.json", err)
	}
}
