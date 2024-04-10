package model

import (
	"encoding/json"
	"log"
	"os"
)

type Tree struct {
	dataPath string
	Folders  []Folder
}

func NewTree() Tree {
	return Tree{
		dataPath: "./data",
	}
}

// Parse returns a slice of all folders in the data path and their respective
// albums
func (t *Tree) Parse() {
	entries, _ := os.ReadDir(t.dataPath)
	for _, e1 := range entries {
		if e1.IsDir() {

			folder := Folder{
				Name: e1.Name(),
			}

			basePath := t.dataPath + "/" + folder.Name

			entries, _ = os.ReadDir(basePath)
			for _, e2 := range entries {
				albumDataRaw, err := os.ReadFile(basePath + "/" + e2.Name() + "/data.json")
				if err != nil {
					log.Println("Coudn't read album data file", err)
				}

				album := Album{}
				err = json.Unmarshal(albumDataRaw, &album)
				if err != nil {
					log.Println("Coudn't parse album data file", err)
				}
				folder.Albums = append(folder.Albums, album)
			}

			t.Folders = append(t.Folders, folder)
		}
	}
}
