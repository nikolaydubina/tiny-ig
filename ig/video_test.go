package ig_test

import (
	"os"
	"testing"

	"github.com/nikolaydubina/tiny-ig/ig"
)

func TestVideo(t *testing.T) {
	exp, err := os.ReadFile("testdata/video.html")
	if err != nil {
		t.Error(err)
	}

	c := ig.Video{Width: 100, Height: 200, Path: "something/video.mov"}
	if v := c.Render(); v != string(exp) {
		t.Errorf("wrong content: %s != %s", v, exp)
	}
}
