package main

import (
	"bytes"
	"errors"
	"io/ioutil"
)

func readImage(imagePath string) []byte {
	data, err := ioutil.ReadFile(imagePath)
	if err != nil {
		panic(err)
	}
	return data
}

func isExif(data []byte) bool {
	signature := data[0:4]
	res := bytes.Compare(signature, []byte("\xFF\xD8\xFF\xE1"))
	if res == 0 {
		return true
	} else {
		return false
	}
}

func findLivePhotoIdx(data []byte) (int, error) {
	if !isExif(data) {
		return -1, errors.New("Not a JPG/EXIF format.")
	}

	// Todo find where mp4 starts

	idx := bytes.Index(data, []byte("\x66\x74\x79\x70"))

	return idx - 4, nil
}

func extractLivePhoto(data []byte, idx int) ([]byte, []byte, error) {
	photo := data[0:idx]
	video := data[idx:]
	return photo, video, nil
}

func main() {
	file := "test.jpg"
	imageData := readImage(file)

	idx, err := findLivePhotoIdx(imageData)
	if err != nil {
		panic(err)
	}

	_, _, _ = extractLivePhoto(imageData, idx)
	// save and exit
}
