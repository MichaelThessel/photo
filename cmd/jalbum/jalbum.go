package main

import (
	"log"
	"os"
	"slices"
	"strings"

	"github.com/MichaelThessel/photo/model"
	"github.com/MichaelThessel/photo/service"

	"golang.org/x/net/html"
)

func main() {
	ap := AlbumParser{
		excluded: []string{
			"hi-res",
			".jalbum",
		},
		importPath: "./import",
		exportPath: "./data",
		thumbnailGenerator: &service.ThumbnailGenerator{
			Size: 400,
		},
	}

	ap.ParseDir()
}

type AlbumParser struct {
	excluded           []string
	importPath         string
	exportPath         string
	thumbnailGenerator *service.ThumbnailGenerator
}

// ParseDir parses the album directory tree and finds almbum subdirectories
func (ap AlbumParser) ParseDir() {
	entries, _ := os.ReadDir(ap.importPath)
	for _, e1 := range entries {
		if e1.IsDir() {

			log.Println("Processing: ", e1.Name())
			folder := &model.Folder{
				Name: e1.Name(),
			}
			folder.CreateFolderDirectory(ap.exportPath)

			basePath := ap.exportPath + "/" + e1.Name()
			entries, _ = os.ReadDir(ap.importPath + "/" + e1.Name())
			for _, e2 := range entries {
				if e2.IsDir() && !slices.Contains(ap.excluded, e2.Name()) {
					albumPath := e1.Name() + "/" + e2.Name()
					log.Println("Processing: ", albumPath)

					album := &model.Album{
						Name: e2.Name(),
					}
					album.CreateAlbumDirectories(basePath)

					// Parse album index files
					indexPath := ap.importPath + "/" + albumPath + "/index.html"
					ap.ParseIndex(indexPath, album)

					// Create the album data file
					album.CreateDataJSON(basePath)

					// Copy the images
					ap.CopyImages(album, folder)
				}
			}
		}
	}
}

// CopyImages copies the images from the JAlbum slide directory to the
// appropriate place in the new folder structure
func (ap AlbumParser) CopyImages(album *model.Album, folder *model.Folder) {
	for _, s := range album.Slides {
		r, err := os.Open(
			ap.importPath + "/" + folder.Name + "/" + album.Name + "/slides/" + s.ImagePath,
		)
		if err != nil {
			log.Println("Couldn't open image source file", err)
			return
		}
		defer r.Close()

		w, err := os.Create(
			ap.exportPath + "/" + folder.Name + "/" + album.Name + "/slides/" + s.ImagePath,
		)
		if err != nil {
			log.Println("Couldn't create image destination file", err)
			return
		}
		defer w.Close()

		_, err = w.ReadFrom(r)
		if err != nil {
			log.Println("Couldn't write image destination file", err)
		}

		ap.thumbnailGenerator.GenerateThumbnail(
			w,
			ap.exportPath+"/"+folder.Name+"/"+album.Name+"/thumbs/"+s.ImagePath,
			false,
		)
	}
}

// ParseIndex parses the index.html files from the albums and extracts the slide
// location and order
func (ap AlbumParser) ParseIndex(path string, album *model.Album) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Could not open album index: ", path)
	}

	indexDoc, err := html.Parse(file)
	if err != nil {
		log.Fatal("Failed to parse index document: ", path)
	}

	// Find slide information
	var parseSlides func(*html.Node)
	parseSlides = func(n *html.Node) {

		if n.Type == html.ElementNode && n.Data == "a" {
			if n.Attr[0].Key == "href" && strings.HasPrefix(n.Attr[0].Val, "slides/") {
				slide := model.Slide{}
				slide.ImagePath = strings.TrimPrefix(n.Attr[0].Val, "slides/")
				if n.NextSibling != nil && n.NextSibling.LastChild != nil && n.NextSibling.LastChild.LastChild != nil {
					slide.Description = n.NextSibling.LastChild.LastChild.Data
				}

				album.Slides = append(album.Slides, slide)
			}

		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parseSlides(c)
		}
	}
	parseSlides(indexDoc)
}
