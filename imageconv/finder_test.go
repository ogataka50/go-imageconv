package imageconv_test

import (
	"testing"

	"errors"
	"reflect"

	"github.com/ogataka50/go-imageconv/imageconv"
)

func TestFinder_IsDir_Success(t *testing.T) {
	testData := []struct {
		dir    string
		ext    string
		expect bool
	}{
		{"../testdata", "jpg", true},
		{"../imageconv", "jpg", true},
		{"../../go-imageconv", "jpg", true},
	}

	for _, d := range testData {
		f := imageconv.Finder{
			Dir: d.dir,
			Ext: d.ext,
		}

		expect := d.expect
		actual, err := f.IsDir()
		if actual != expect || err != nil {
			t.Errorf(`expect="%v" actual="%v"`, expect, actual)
		}
	}
}

func TestFinder_IsDir_Error(t *testing.T) {
	testData := []struct {
		dir    string
		ext    string
		expect bool
		err    error
	}{
		{"../xxx", "jpg", false, errors.New("stat ../xxx: no such file or directory")},
		{"../testdata/sample.png", "png", false, nil},
	}

	for _, d := range testData {
		f := imageconv.Finder{
			Dir: d.dir,
			Ext: d.ext,
		}

		expect := d.expect
		actual, err := f.IsDir()
		if actual != expect || (err != nil && err.Error() != d.err.Error()) {
			t.Errorf(`expect="%v" actual="%v" err="%v"`, expect, actual, err)
		}
	}
}

func TestFinder_FindByExt_Success(t *testing.T) {
	testData := []struct {
		dir        string
		ext        string
		expectList []string
	}{
		{"../testdata", "png", []string{"../testdata/sample.png"}},
	}

	for _, d := range testData {
		f := imageconv.Finder{
			Dir: d.dir,
			Ext: d.ext,
		}

		expect := d.expectList
		actual, err := f.FindByExt()
		if !reflect.DeepEqual(expect, actual) || err != nil {
			t.Errorf(`expect="%v" actual="%v" err="%v"`, expect, actual, err)
		}
	}
}
