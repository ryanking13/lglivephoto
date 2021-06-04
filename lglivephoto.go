package lglivephoto // import "github.com/ryanking13/lglivephoto"

import (
	"io/ioutil"

	"go.uber.org/zap/zapcore"
)

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

func Pack(imagePath string, videoPath string) ([]byte, error) {
	Logger.Debugf("Start packing: %s to %s", videoPath, imagePath)
	return nil, nil
}

func Debug(debug bool) {
	Atom.SetLevel(zapcore.DebugLevel)
}
