package lglivephoto_test

import (
	"fmt"
	"testing"

	"github.com/ryanking13/lglivephoto"
)

func TestUnpack(t *testing.T) {
	image, video, err := lglivephoto.Unpack("test_images/test.jpg")
	fmt.Println(image, video, err)
	lglivephoto.Pack("1", "2")
}
