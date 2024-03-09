package util

import (
	"bytes"
	"io"
	"log"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func ShiftJISToUTF8(sjisStr []byte) []byte {
	reader := transform.NewReader(bytes.NewReader(sjisStr), japanese.ShiftJIS.NewDecoder())
	utf8Str, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	return utf8Str
}
