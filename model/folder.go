package model

import (
	"log"
	"os"
)

type Folder struct {
	Name      string
	ThumbPath string
	Albums    []Album
}

// CreateFolderDirectory creates directories for base folders
func (f Folder) CreateFolderDirectory(basePath string) {
	// Create album directories
	err := os.Mkdir(basePath+"/"+f.Name, 0744)
	if err != nil {
		log.Println("Error creating folder directory", err)
	}
}
