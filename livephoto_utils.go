package lglivephoto

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// Referenced from official golang JPEG implementation: https://github.com/golang/go/tree/master/src/image/jpeg
const (
	sof0Marker = '\xc0' // Start Of Frame (Baseline Sequential).
	sof1Marker = '\xc1' // Start Of Frame (Extended Sequential).
	sof2Marker = '\xc2' // Start Of Frame (Progressive).
	dhtMarker  = '\xc4' // Define Huffman Table.
	rst0Marker = '\xd0' // ReSTart (0).
	rst7Marker = '\xd7' // ReSTart (7).
	soiMarker  = '\xd8' // Start Of Image.
	eoiMarker  = '\xd9' // End Of Image.
	sosMarker  = '\xda' // Start Of Scan.
	dqtMarker  = '\xdb' // Define Quantization Table.
	driMarker  = '\xdd' // Define Restart Interval.
	comMarker  = '\xfe' // COMment.
	// "APPlication specific" markers aren't part of the JPEG spec per se,
	// but in practice, their use is described at
	// https://www.sno.phy.queensu.ca/~phil/exiftool/TagNames/JPEG.html
	app0Marker  = '\xe0'
	app1Marker  = '\xe1'
	app2Marker  = '\xe2'
	app14Marker = '\xee'
	app15Marker = '\xef'
)

var jpegMarkers = map[byte]bool{
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
	return bytes.Compare(signature, []byte("\xFF\xD8")) == 0
}

func isSOS(signature []byte) bool {
	return bytes.Compare(signature, []byte("\xFF\xDA")) == 0
}

func isEOI(signature []byte) bool {
	return bytes.Compare(signature, []byte("\xFF\xD9")) == 0
}

func isMP4(signature []byte) bool {
	return bytes.Compare(signature, []byte("\x66\x74\x79\x70")) == 0
}

func isJPEGMarker(signature []byte) bool {
	if len(signature) < 2 {
		return false
	}

	_, jpegMarker := jpegMarkers[signature[1]]
	return signature[0] == '\xFF' && jpegMarker
}

func findVideoIndex(data []byte) (int, error) {

	if !isJPEG(data[0:2]) {
		return -1, errors.New("Not a JPEG format.")
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
		return -1, errors.New("No Embedded video found")
	}

	if !isMP4(data[videoStartIdx : videoStartIdx+4]) {
		return -1, errors.New(`There is a chunck after JPEG image, but it is not a MP4 file, maybe the image is corrupted.
							  Please report the problem to: https://github.com/ryanking13/lglivephoto/issues`)
	}

	return videoStartIdx - 4, nil
}

func extractLivePhoto(data []byte, idx int) ([]byte, []byte, error) {
	photo := data[0:idx]
	video := data[idx:]
	return photo, video, nil
}
