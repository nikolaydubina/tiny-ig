package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"text/template"

	_ "embed"
)

//go:embed gallery.html
var galleryHTMLTemplate string

type PhotoParams struct {
	Path string
}

type GalleryParams struct {
	Photos []PhotoParams
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

	params := GalleryParams{
		Photos: []PhotoParams{},
	}

	files, err := os.ReadDir(photosDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		params.Photos = append([]PhotoParams{{
			Path: strings.Join([]string{"", photosDir, file.Name()}, "/"),
		}}, params.Photos...)
	}

	parsedTemplate, err := template.New("gallery").Parse(galleryHTMLTemplate)
	if err != nil {
		log.Fatal(err)
	}
	if err := parsedTemplate.Execute(os.Stdout, params); err != nil {
		log.Fatal(err)
	}
}
