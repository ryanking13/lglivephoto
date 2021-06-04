// Package lglivephoto provides functions to unpack or pack LG Live Photo.
// LG Live Photo is a functionality in LG smartphones
// where actions before and after taking a photo is recorded with the photo.
// (see: https://www.lg.com/uk/support/product-help/CT00008356-20150844039308)
package lglivephoto // import "github.com/ryanking13/lglivephoto"

import (
	"io/ioutil"

	"go.uber.org/zap/zapcore"
)

// Unpack takes a Live Photo file and extracts it to
// image and video embedded in the photo.
func Unpack(imagePath string) ([]byte, []byte, error) {
	Logger.Debugf("Start unpacking: %s", imagePath)
	data, err := ioutil.ReadFile(imagePath)

	if err != nil {
		return nil, nil, err
	}

	idx, err := findVideoIndex(data)

	if err != nil {
		return nil, nil, err
	}

	Logger.Debugf("Video index of %s = %x", imagePath, idx)
	return data[:idx], data[idx:], nil
}

// Pack generates LG Live Photo by embedding the mp4 video to the image,
func Pack(imagePath string, videoPath string) ([]byte, error) {
	Logger.Debugf("Start packing: %s to %s", videoPath, imagePath)
	return nil, nil
}

func Debug(debug bool) {
	Atom.SetLevel(zapcore.DebugLevel)
}
