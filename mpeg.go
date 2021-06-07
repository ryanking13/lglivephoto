package lglivephoto

import (
	"bytes"
)

func isMP4(signature []byte) bool {
	return bytes.Equal(signature[0:3], []byte("\x00\x00\x00")) && bytes.Equal(signature[4:8], []byte("\x66\x74\x79\x70"))
}
