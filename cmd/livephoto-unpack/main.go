package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ryanking13/lglivephoto"
)

func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), nil
}

func main() {
	target := flag.String("target", "", "Target livephoto file or directory which contains livephoto files")
	verbose := flag.Bool("verbose", false, "Print logs (used for debugging)")

	flag.Parse()
	lglivephoto.Debug(*verbose)

	isDir, err := isDirectory(*target)

	if err != nil {
		panic(err)
	}

	var targets []string

	if isDir {
		targetsPath := fmt.Sprintf("%s/*.jpg", *target)
		targets, err = filepath.Glob(targetsPath)

		if err != nil {
			panic(err)
		}

	} else {
		targets = []string{*target}
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
