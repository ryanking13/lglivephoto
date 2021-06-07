// Package lglivephoto provides functions to unpack or pack LG Live Photo.
// LG Live Photo is a functionality in LG smartphones
// where actions before and after taking a photo is recorded with the photo.
// (see: https://www.lg.com/uk/support/product-help/CT00008356-20150844039308)
package lglivephoto // import "github.com/ryanking13/lglivephoto"

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"

	"go.uber.org/zap/zapcore"
)

// Unpack takes a Live Photo file and extracts it to
// image and video embedded in the photo.
func Unpack(imagePath string) ([]byte, []byte, error) {
	logger.Debugf("Start unpacking: %s", imagePath)
	data, err := ioutil.ReadFile(imagePath)

	if err != nil {
		return nil, nil, err
	}

	idx, err := findVideoIndex(data)

	if err != nil {
		return nil, nil, err
	}

	logger.Debugf("Video index of %s = %x", imagePath, idx)
	return data[:idx], data[idx:], nil
}

// Pack generates LG Live Photo by embedding the mp4 video to the image,
func Pack(imagePath string, videoPath string) ([]byte, error) {
	logger.Debugf("Start packing: %s to %s", videoPath, imagePath)
	imageData, err := ioutil.ReadFile(imagePath)

	if err != nil {
		return nil, err
	}

	videoData, err := ioutil.ReadFile(videoPath)

	if err != nil {
		return nil, err
	}

	if !isJPEG(imageData[0:2]) {
		return nil, fmt.Errorf("%s is not a JPEG file", imagePath)
	}

	if !isMP4(videoData[0:8]) {
		return nil, fmt.Errorf("%s is not a MP4 file", videoPath)
	}

	_, err = findVideoIndex(imageData)
	if err == nil {
		return nil, fmt.Errorf("%s is already a Live Photo", imagePath)
	}

	return append(imageData, videoData...), nil
}

// IsLivePhoto checks whether given image file is a live photo or not.
func IsLivePhoto(imagePath string) (bool, error) {
	logger.Debugf("Checking whether %s is a live photo", imagePath)
	data, err := ioutil.ReadFile(imagePath)

	if err != nil {
		return false, err
	}

	_, err = findVideoIndex(data)

	if err != nil {
		return false, err
	}

	return true, nil
}

// Debug changes logger verbosity, which is mostly used for debugging.
func Debug(debug bool) {
	if debug {
		atom.SetLevel(zapcore.DebugLevel)
	} else {
		atom.SetLevel(zapcore.InfoLevel)
	}
}

func findVideoIndex(data []byte) (int, error) {

	if !isJPEG(data[0:2]) {
		return -1, errors.New("not a JPEG format")
	}

	appIdx := 2
	for isJPEGMarker(data[appIdx:appIdx+2]) && !isSOS(data[appIdx:appIdx+2]) {
		segmentSize := binary.BigEndian.Uint16(data[appIdx+2 : appIdx+4])
		appIdx += int(segmentSize) + 2

		// fmt.Printf("%x %x %x\n", segmentSize, appIdx, data[appIdx:appIdx+2])
	}

	// https://stackoverflow.com/questions/26715684/parsing-jpeg-sos-marker
	eoiIdx := appIdx + 2
	for !isEOI(data[eoiIdx : eoiIdx+2]) {
		eoiIdx++
	}

	videoStartIdx := eoiIdx + 2
	if videoStartIdx >= len(data) {
		return -1, errors.New("not a live photo")
	}

	if !isMP4(data[videoStartIdx : videoStartIdx+4]) {
		return -1, errors.New(
			"there is a chunck after JPEG image, but it is not a MP4 file, maybe the image is corrupted. " +
				"Please report the problem to: https://github.com/ryanking13/lglivephoto/issues",
		)
	}

	return videoStartIdx, nil
}
