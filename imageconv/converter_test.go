package imageconv_test

import (
	"os"
	"testing"

	"github.com/ogataka50/go-imageconv/imageconv"
)

func TestConverter_Convert(t *testing.T) {
	cases := []struct {
		name    string
		path    string
		fromExt string
		toExt   string
	}{
		{name: "pngTojpg", path: "../testdata/sample_png.png", fromExt: "png", toExt: "jpg"},
		{name: "pngTogif", path: "../testdata/sample_png.png", fromExt: "png", toExt: "gif"},
		{name: "gifTopng", path: "../testdata/sample_gif.gif", fromExt: "gif", toExt: "png"},
		{name: "gifTojpg", path: "../testdata/sample_gif.gif", fromExt: "gif", toExt: "jpg"},
		{name: "jpgTopng", path: "../testdata/sample_jpg.jpg", fromExt: "jpg", toExt: "png"},
		{name: "jpgTogif", path: "../testdata/sample_jpg.jpg", fromExt: "jpg", toExt: "gif"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			conv := imageconv.Converter{
				Path:    c.path,
				FromExt: c.fromExt,
				ToExt:   c.toExt,
			}
			testConvert(t, conv)
		})
	}
}

func testConvert(t *testing.T, c imageconv.Converter) {
	t.Helper()

	convertedPath, err := c.Convert()
	if err != nil {
		t.Errorf(`Converter="%v"`, c)
	}
	defer os.Remove(convertedPath)
}
