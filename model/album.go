package model

import (
	"bytes"
	"encoding/json"
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
