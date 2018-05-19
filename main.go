package main

import (
	"flag"
	"sync"

	"github.com/ogataka50/go-imageconv/imageconv"
)

var fromExt string
var toExt string
var dir string

func init() {
	flag.StringVar(&fromExt, "from", "jpg", "target ext")
	flag.StringVar(&toExt, "to", "png", "converted ext")
	flag.StringVar(&dir, "dir", "", "target dir")
}

func main() {

	flag.Parse()

	//TODO 対象extかチェック
	if fromExt == toExt {
		panic("Invalid ext : from => " + fromExt + ", to => " + toExt)
	}

	f := imageconv.Finder{
		Dir: dir,
		Ext: fromExt,
	}

	// chk dir exists
	if !f.IsDir() {
		panic("dir not exists : " + f.Dir)
	}

	//file finder
	fList, err := f.FindByExt()
	if err != nil {
		panic(err)
	}

	//img convert
	wg := &sync.WaitGroup{}
	for _, v := range fList {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c := imageconv.Converter{
				Path:    v,
				FromExt: fromExt,
				ToExt:   toExt,
			}
			c.Convert()
		}()
	}
	wg.Wait()

}
