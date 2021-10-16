package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Photo struct {
	Path   string
	Width  int
	Height int
}

func (p Photo) Render() string {
	s := ""
	s += fmt.Sprintf(`<a href="%s" target="_blank" style="display: contents;">`, p.Path) + "\n"
	s += fmt.Sprintf(`<img src="%s" style="margin: 8px; object-fit: cover; width: %dpx; height: %dpx;" loading="lazy">`, p.Path, p.Width, p.Height) + "\n"
	s += "</a>\n"
	return s
}

type Gallery struct {
	Photos []Photo
}

func (g Gallery) Render() string {
	if len(g.Photos) == 0 {
		return ""
	}

	s := ""
	s += `<div align="center">` + "\n"
	for _, p := range g.Photos {
		s += p.Render() + "\n"
	}
	s += "</div>\n"

	return s
}

func main() {
	var (
		photosDir string
	)

	flag.StringVar(&photosDir, "dir", "", "path to directory with photos")
	flag.Parse()

	if photosDir == "" {
		log.Fatal("dir arg is missing")
	}

	gallery := Gallery{}

	files, err := os.ReadDir(photosDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		path := strings.Join([]string{"", photosDir, file.Name()}, "/")
		photo := Photo{
			Path:   path,
			Width:  292,
			Height: 292,
		}
		gallery.Photos = append(gallery.Photos, photo)
	}

	sort.Slice(gallery.Photos, func(i, j int) bool { return gallery.Photos[i].Path > gallery.Photos[j].Path })

	os.Stdout.WriteString(gallery.Render())
}
