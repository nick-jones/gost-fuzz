package gost_fuzz

import (
	"bytes"

	"github.com/nick-jones/gost/pkg/scan"
)

func Fuzz(data []byte) int {
	r := bytes.NewReader(data)
	if _, err := scan.Run(r); err != nil {
		return 0
	}
	return 1
}
