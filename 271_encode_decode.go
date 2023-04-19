package main

import (
	"bytes"
)

type Codec struct {
}

func (codec *Codec) Encode(strs []string) string {
	var b bytes.Buffer
	return b.String()
}

func (codec *Codec) Decode(strs string) []string {
	s := make([]string, 0)

	return s
}
