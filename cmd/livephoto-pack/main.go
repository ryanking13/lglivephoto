package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	flags "github.com/jessevdk/go-flags"
	"github.com/ryanking13/lglivephoto"
)

var opts struct {
	Verbose bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	Targets struct {
		Image string `positional-arg-name:"IMAGE" description:"Template image file which video will be embedded"`
		Video string `positional-arg-name:"VIDEO" description:"Video file which will be embedded to the template Image"`
	} `positional-args:"yes" required:"yes"`
}

func pack(image string, video string) {

	livephoto, err := lglivephoto.Pack(image, video)
	if err != nil {
		fmt.Printf("[-] Fail (%s / %s): %s\n", image, video, err.Error())
		return
	}

	savePath := strings.TrimSuffix(image, filepath.Ext(image)) + "_livephoto.jpg"
	err = ioutil.WriteFile(savePath, livephoto, 0644)
	if err != nil {
		fmt.Printf("[-] Fail (%s / %s): %s\n", image, video, err.Error())
		return
	}
}

func main() {

	if _, err := flags.Parse(&opts); err != nil {
		switch flagsErr := err.(type) {
		case *flags.Error:
			if flagsErr.Type == flags.ErrHelp {
				os.Exit(0)
			} else {
				// fmt.Println(err.Error())
				os.Exit(1)
			}
		default:
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	lglivephoto.Debug(opts.Verbose)
	pack(opts.Targets.Image, opts.Targets.Video)
}
