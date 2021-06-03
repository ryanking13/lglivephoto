package lglivephoto // import "github.com/ryanking13/lglivephoto"

import (
	"io/ioutil"
)

func Unpack(imagePath string) ([]byte, []byte, error) {
	data, err := ioutil.ReadFile(imagePath)

	if err != nil {
		return nil, nil, err
	}

	idx, err := findVideoIndex(data)

	if err != nil {
		return nil, nil, err
	}

	return data[:idx], data[idx:], nil
}

func Pack(imagePath string, videoPath string) ([]byte, error) {
	return nil, nil
}
