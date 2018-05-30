package imageconv_test

import (
	"os"
	"testing"

	"github.com/ogataka50/go-imageconv/imageconv"
)

func TestConverter_Convert(t *testing.T) {
	c := imageconv.Converter{
		Path:    "../testdata/sample.png",
		FromExt: "png",
		ToExt:   "jpg",
	}

	convertedPath, err := c.Convert()
	if err != nil {
		t.Errorf(`Converter="%v"`, c)
	}
	defer os.Remove(convertedPath)
}
