package lglivephoto

import (
	"bytes"
	"encoding/binary"
	"errors"
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

func isMP4(signature []byte) bool {
	return bytes.Equal(signature[0:3], []byte("\x00\x00\x00")) && bytes.Equal(signature[4:8], []byte("\x66\x74\x79\x70"))
}

func isJPEGMarker(signature []byte) bool {
	_, exists := jpegMarkers[string(signature)]
	return exists
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
		return -1, errors.New("no embedded video found")
	}

	if !isMP4(data[videoStartIdx : videoStartIdx+4]) {
		return -1, errors.New(
			"there is a chunck after JPEG image, but it is not a MP4 file, maybe the image is corrupted. " +
				"Please report the problem to: https://github.com/ryanking13/lglivephoto/issues",
		)
	}

	return videoStartIdx - 4, nil
}
