package ig_test

import (
	"os"
	"testing"

	"github.com/nikolaydubina/tiny-ig/ig"
)

func TestPhoto(t *testing.T) {
	exp, err := os.ReadFile("testdata/photo.html")
	if err != nil {
		t.Error(err)
	}

	c := ig.Photo{Width: 100, Height: 200, Path: "something/photo.jpeg"}
	if v := c.Render(); v != string(exp) {
		t.Errorf("wrong content: %s != %s", v, exp)
	}
}
