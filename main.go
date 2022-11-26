package main

import (
	"flag"
	"log"
	"os"
	"sort"

	"github.com/nikolaydubina/tiny-ig/ext"
	"github.com/nikolaydubina/tiny-ig/ig"
)

func main() {
	var (
		photosDir string
		prefix    string
		width     int
		height    int
	)

	flag.StringVar(&photosDir, "dir", ".", "path to directory")
	flag.StringVar(&prefix, "prefix", "", "prefix for URLs")
	flag.IntVar(&width, "w", 292, "width of post")
	flag.IntVar(&height, "h", 292, "heigh of post")
	flag.Parse()

	gallery := ig.Gallery{}

	files, err := os.ReadDir(photosDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		switch {
		case ext.IsImage(file.Name()):
			photo := ig.Photo{
				Path:   prefix + file.Name(),
				Width:  width,
				Height: height,
			}
			gallery.Posts = append(gallery.Posts, photo)
		case ext.IsVideo(file.Name()):
			video := ig.Video{
				Path:   prefix + file.Name(),
				Width:  width,
				Height: height,
			}
			gallery.Posts = append(gallery.Posts, video)
		default:
			continue
		}
	}

	sort.Slice(gallery.Posts, func(i, j int) bool { return gallery.Posts[i].Date() > gallery.Posts[j].Date() })

	gallery.RenderInto(os.Stdout)
}
