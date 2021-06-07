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
		Args []string `positional-arg-name:"TARGET" description:"Target livephoto file or directory which contains livephoto files"`
	} `positional-args:"yes" required:"yes"`
}

func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), nil
}

func unpack(target string) {
	var targets []string

	isDir, err := isDirectory(target)

	if err != nil {
		panic(err)
	}

	if isDir {
		targetsPath := fmt.Sprintf("%s/*.jpg", target)
		targets, err = filepath.Glob(targetsPath)

		if err != nil {
			panic(err)
		}

	} else {
		targets = []string{target}
	}

	for _, targetImage := range targets {
		image, video, err := lglivephoto.Unpack(targetImage)
		if err != nil {
			fmt.Printf("[-] Fail (%s): %s\n", targetImage, err.Error())
		}

		err = ioutil.WriteFile(strings.TrimSuffix(targetImage, filepath.Ext(targetImage))+"_unpack.jpg", image, 0644)
		if err != nil {
			fmt.Printf("[-] Fail (%s): %s\n", targetImage, err.Error())
		}

		err = ioutil.WriteFile(strings.TrimSuffix(targetImage, filepath.Ext(targetImage))+"_unpack.mp4", video, 0644)
		if err != nil {
			fmt.Printf("[-] Fail (%s): %s\n", targetImage, err.Error())
		}
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

	if len(opts.Targets.Args) == 0 {
		fmt.Println("File or directory not specified, use --help to see usage")
		os.Exit(1)
	}

	lglivephoto.Debug(opts.Verbose)

	for _, target := range opts.Targets.Args {
		unpack(target)
	}
}
