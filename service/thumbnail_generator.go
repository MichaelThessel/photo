package service

import (
	"log"
	"os"

	"github.com/disintegration/imaging"
)

type ThumbnailGenerator struct {
	Size int
}

func (tg *ThumbnailGenerator) GenerateAlbumThumbnail(file *os.File, outputPath string, force bool) {
	// Don't generate if thumbnail already exists
	if !force {
		if _, err := os.Stat(outputPath); err == nil {
			return
		}
	}

	log.Println("Generating album thumbnail", outputPath)

	thumb, err := imaging.Open(file.Name())
	if err != nil {
		log.Fatalln("Couldn't open source image", err)
	}
	thumb = imaging.Resize(thumb, tg.Size, 0, imaging.Lanczos)

	err = imaging.Save(thumb, outputPath)
	if err != nil {
		log.Fatalln("Couldn't save thumbnail", err)
	}
}

func (tg *ThumbnailGenerator) GenerateFolderThumbnail(file *os.File, outputPath string, force bool) {
	// Don't generate if thumbnail already exists
	if !force {
		if _, err := os.Stat(outputPath); err == nil {
			return
		}
	}

	log.Println("Generating folder thumbnail", outputPath)

	thumb, err := imaging.Open(file.Name())
	if err != nil {
		log.Fatalln("Couldn't open source image", err)
	}

	// Convert to square
	bounds := thumb.Bounds()
	if bounds.Dx() > bounds.Dy() {
		thumb = imaging.CropCenter(thumb, bounds.Dy(), bounds.Dy())
	} else {
		thumb = imaging.CropCenter(thumb, bounds.Dx(), bounds.Dx())
	}
	thumb = imaging.Resize(thumb, tg.Size, 0, imaging.Lanczos)

	err = imaging.Save(thumb, outputPath)
	if err != nil {
		log.Fatalln("Couldn't save thumbnail", err)
	}
}
