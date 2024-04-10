package service

import (
	"log"
	"os"

	"github.com/disintegration/imaging"
)

type ThumbnailGenerator struct {
	Size int
}

func (tg *ThumbnailGenerator) GenerateThumbnail(file *os.File, outputPath string, force bool) {
	// Don't generate if thumbnail already exists
	if !force {
		if _, err := os.Stat(outputPath); err == nil {
			return
		}
	}

	log.Println("Generating Thumbnail", outputPath)

	image, err := imaging.Open(file.Name())
	if err != nil {
		log.Fatalln("Couldn't open source image", err)
	}
	thumb := imaging.Resize(image, tg.Size, 0, imaging.Lanczos)

	err = imaging.Save(thumb, outputPath)
	if err != nil {
		log.Fatalln("Couldn't save thumbnail", err)
	}
}
