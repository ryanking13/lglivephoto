package lglivephoto

import (
	"bytes"
	"image"
)

// Referenced from official golang JPEG implementation: https://github.com/golang/go/tree/master/src/image/jpeg
var (
	sof0Marker = "\xff\xc0" // Start Of Frame (Baseline Sequential).
	sof1Marker = "\xff\xc1" // Start Of Frame (Extended Sequential).
	sof2Marker = "\xff\xc2" // Start Of Frame (Progressive).
	dhtMarker  = "\xff\xc4" // Define Huffman Table.
	rst0Marker = "\xff\xd0" // ReSTart (0).
	rst7Marker = "\xff\xd7" // ReSTart (7).
	soiMarker  = "\xff\xd8" // Start Of Image.
	eoiMarker  = "\xff\xd9" // End Of Image.
	sosMarker  = "\xff\xda" // Start Of Scan.
	dqtMarker  = "\xff\xdb" // Define Quantization Table.
	driMarker  = "\xff\xdd" // Define Restart Interval.
	comMarker  = "\xff\xfe" // COMment.
	// "APPlication specific" markers aren't part of the JPEG spec per se,
	// but in practice, their use is described at
	// https://www.sno.phy.queensu.ca/~phil/exiftool/TagNames/JPEG.html
	app0Marker  = "\xff\xe0"
	app1Marker  = "\xff\xe1"
	app2Marker  = "\xff\xe2"
	app14Marker = "\xff\xee"
	app15Marker = "\xff\xef"
)

var jpegMarkers = map[string]bool{
	sof0Marker:  true,
	sof1Marker:  true,
	sof2Marker:  true,
	dhtMarker:   true,
	rst0Marker:  true,
	rst7Marker:  true,
	soiMarker:   true,
	eoiMarker:   true,
	sosMarker:   true,
	dqtMarker:   true,
	driMarker:   true,
	comMarker:   true,
	app0Marker:  true,
	app1Marker:  true,
	app2Marker:  true,
	app14Marker: true,
	app15Marker: true,
}

func isJPEG(signature []byte) bool {
	return bytes.Equal(signature, []byte("\xFF\xD8"))
}

func isSOS(signature []byte) bool {
	return bytes.Equal(signature, []byte("\xFF\xDA"))
}

func isEOI(signature []byte) bool {
	return bytes.Equal(signature, []byte("\xFF\xD9"))
}

func isJPEGMarker(signature []byte) bool {
	_, exists := jpegMarkers[string(signature)]
	return exists
}

func jpegSize(data []byte) (int, int, error) {
	reader := bytes.NewReader(data)

	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		return -1, -1, err
	}

	return img.Width, img.Height, nil
}
