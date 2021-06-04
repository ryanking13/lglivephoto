package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

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

	flag.Parse()

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

		err = ioutil.WriteFile("", image, 0644)
		if err != nil {
			fmt.Printf("[-] Fail (%s): %s\n", targetImage, err.Error())
		}

		err = ioutil.WriteFile("", video, 0644)
		if err != nil {
			fmt.Printf("[-] Fail (%s): %s\n", targetImage, err.Error())
		}
	}
}
