/*
 Image Converter

 Arguments...
  from	convert from ext
  to	convert to ext
  dir	target convert dir
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/ogataka50/go-imageconv/imageconv"
)

var (
	fromExt string
	toExt   string
	dir     string
)

func init() {
	flag.StringVar(&fromExt, "from", "jpg", "target ext")
	flag.StringVar(&toExt, "to", "png", "converted ext")
	flag.StringVar(&dir, "dir", "", "target dir")
}

func main() {

	flag.Parse()

	//TODO Check support ext
	if fromExt == toExt {
		fmt.Fprintf(os.Stderr, "Invalid ext : from => %s, to => %s\n", fromExt, toExt)
		os.Exit(1)

	}

	f := imageconv.Finder{
		Dir: dir,
		Ext: fromExt,
	}

	// chk dir exists
	isDir, err := f.IsDir()
	if !isDir || err != nil {
		fmt.Fprintf(os.Stderr, "Dir not exists : %s\n", f.Dir)
		os.Exit(1)
	}

	//file finder
	fList, err := f.FindByExt()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error : %v\n", err)
		os.Exit(1)
	}

	//img convert
	wg := &sync.WaitGroup{}
	for _, v := range fList {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			c := imageconv.Converter{
				Path:    path,
				FromExt: fromExt,
				ToExt:   toExt,
			}
			c.Convert()
		}(v)
	}
	wg.Wait()

	os.Exit(0)
}
