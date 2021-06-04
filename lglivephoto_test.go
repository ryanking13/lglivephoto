package lglivephoto_test

import (
	"testing"

	"github.com/ryanking13/lglivephoto"
)

func TestUnpackSuccess(t *testing.T) {
	_, _, err := lglivephoto.Unpack("test_images/test.jpg")
	if err != nil {
		t.Error(err)
	}
}

func TestPack(t *testing.T) {
	// TODO
	// lglivephoto.Pack("1", "2")
}
