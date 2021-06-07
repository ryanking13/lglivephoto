package lglivephoto_test

import (
	"testing"

	"github.com/ryanking13/lglivephoto"
)

func TestUnpackSuccess(t *testing.T) {
	_, _, err := lglivephoto.Unpack("test_images/livephoto.jpg")
	if err != nil {
		t.Error(err.Error())
	}
}

func TestUnpackFail(t *testing.T) {
	_, _, err := lglivephoto.Unpack("test_images/non_livephoto.jpg")
	if err == nil {
		t.Error("Didn't fail unpacking when trying to unpack non-livephoto")
	}
}

func TestPackSuccess(t *testing.T) {
	_, err := lglivephoto.Pack("test_images/non_livephoto.jpg", "test_images/test_video.mp4")
	if err != nil {
		t.Error(err.Error())
	}
}

func TestPackFail(t *testing.T) {
	_, err := lglivephoto.Pack("test_images/livephoto.jpg", "test_images/test_video.mp4")
	if err == nil {
		t.Error("Didn't fail packing when trying to embed a video to a live photo")
	}
}
