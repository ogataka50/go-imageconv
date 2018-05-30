package imageconv_test

import (
	"testing"

	"errors"
	"reflect"

	"github.com/ogataka50/go-imageconv/imageconv"
)

func TestFinder_IsDir_Success(t *testing.T) {
	cases := []struct {
		name   string
		dir    string
		ext    string
		expect bool
	}{
		{name: "../testdata", dir: "../testdata", ext: "jpg", expect: true},
		{name: "../imageconv", dir: "../imageconv", ext: "jpg", expect: true},
		{name: "../../go-imageconv", dir: "../../go-imageconv", ext: "jpg", expect: true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f := imageconv.Finder{
				Dir: c.dir,
				Ext: c.ext,
			}

			expect := c.expect
			actual, err := f.IsDir()
			if actual != expect || err != nil {
				t.Errorf(`expect="%v" actual="%v"`, expect, actual)
			}
		})
	}
}

func TestFinder_IsDir_Error(t *testing.T) {
	cases := []struct {
		name   string
		dir    string
		ext    string
		expect bool
		err    error
	}{
		{name: "../xxx", dir: "../xxx", ext: "jpg", err: errors.New("stat ../xxx: no such file or directory")},
		{name: "../xxx", dir: "../testdata/sample_png.png", ext: "png"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f := imageconv.Finder{
				Dir: c.dir,
				Ext: c.ext,
			}

			expect := c.expect
			actual, err := f.IsDir()
			if actual != expect || (err != nil && err.Error() != c.err.Error()) {
				t.Errorf(`expect="%v" actual="%v" err="%v"`, expect, actual, err)
			}
		})
	}
}

func TestFinder_FindByExt_Success(t *testing.T) {
	cases := []struct {
		name       string
		dir        string
		ext        string
		expectList []string
	}{
		{name: "../testdsata", dir: "../testdata", ext: "png", expectList: []string{"../testdata/sample_png.png"}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {

			f := imageconv.Finder{
				Dir: c.dir,
				Ext: c.ext,
			}

			expect := c.expectList
			actual, err := f.FindByExt()
			if !reflect.DeepEqual(expect, actual) || err != nil {
				t.Errorf(`expect="%v" actual="%v" err="%v"`, expect, actual, err)
			}
		})
	}
}
