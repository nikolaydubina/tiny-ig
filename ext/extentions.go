package ext

import "strings"

var imageExtensions = map[string]bool{
	"jpeg": true,
	"jpg":  true,
}

var videoExtensions = map[string]bool{
	"mov": true,
}

func IsImage(filename string) bool {
	parts := strings.Split(filename, ".")
	if len(parts) < 1 {
		return false
	}
	ext := parts[len(parts)-1]
	return imageExtensions[strings.ToLower(ext)]
}

func IsVideo(filename string) bool {
	parts := strings.Split(filename, ".")
	if len(parts) < 1 {
		return false
	}
	ext := parts[len(parts)-1]
	return videoExtensions[strings.ToLower(ext)]
}
