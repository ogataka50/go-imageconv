package imageconv

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

type Converter struct {
	Path    string
	FromExt string
	ToExt   string
}

// Read image file from file path
func (c Converter) read() (image.Image, error) {
	file, err := os.Open(c.Path)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	var img image.Image

	switch c.FromExt {
	case "jpg", "jpeg":
		img, err = jpeg.Decode(file)
	case "gif":
		img, err = gif.Decode(file)
	case "png":
		img, err = png.Decode(file)
	default:
		err = errors.New("Not Support Ext")
	}

	if err != nil {
		return nil, err
	}

	return img, err
}

// Write image file
func (c Converter) write(img image.Image) (string, error) {
	convertedPath := c.Path[:len(c.Path)-len(filepath.Ext(c.Path))] + "." + c.ToExt

	convertedFile, err := os.Create(convertedPath)
	defer convertedFile.Close()

	switch c.ToExt {
	case "jpg", "jpeg":
		jpeg.Encode(convertedFile, img, &jpeg.Options{})
	case "gif":
		gif.Encode(convertedFile, img, &gif.Options{})
	case "png":
		png.Encode(convertedFile, img)
	default:
		err = errors.New("Not Support Ext")
	}

	return convertedPath, err
}

// Convert image file fromExt -> toExt
func (c Converter) Convert() error {
	img, err := c.read()
	if err != nil {
		return err
	}

	convertedPath, err := c.write(img)
	if err != nil {
		return err
	}
	fmt.Println(c.Path + " -> " + convertedPath)

	return nil
}
