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

// image converter
type Converter struct {
	Path    string
	FromExt string
	ToExt   string
}

// Read image file from file path
func (c Converter) read() (image.Image, error) {
	file, err := os.Open(c.Path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var img image.Image

	switch c.FromExt {
	case "jpg", "jpeg":
		img, err = jpeg.Decode(file)
	case "gif":
		img, err = gif.Decode(file)
	case "png":
		img, err = png.Decode(file)
	default:
		err = errors.New("Not Support Ext\n")
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
	if err != nil {
		return "", err
	}
	defer convertedFile.Close()

	switch c.ToExt {
	case "jpg", "jpeg":
		jpeg.Encode(convertedFile, img, &jpeg.Options{})
	case "gif":
		gif.Encode(convertedFile, img, &gif.Options{})
	case "png":
		png.Encode(convertedFile, img)
	default:
		err = errors.New("Not Support Ext\n")
	}

	return convertedPath, err
}

// Convert image file fromExt -> toExt
func (c Converter) Convert() (string, error) {
	img, err := c.read()
	if err != nil {
		return "", err
	}

	convertedPath, err := c.write(img)
	if err != nil {
		return "", err
	}
	fmt.Println(c.Path + " -> " + convertedPath)

	return convertedPath, nil
}
