package imageconv

import (
	"os"
	"path/filepath"
)

type Finder struct {
	Dir string
	Ext string
}

// Check Is Dir
func (f Finder) IsDir() bool {
	fInfo, err := os.Stat(f.Dir)
	if err != nil {
		return false
	}

	return fInfo.IsDir()
}

// Find files by Ext recursively
func (f Finder) FindByExt() ([]string, error) {
	var fList []string

	err := filepath.Walk(f.Dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == "."+f.Ext {
			fList = append(fList, path)
		}

		return nil
	})

	return fList, err
}
