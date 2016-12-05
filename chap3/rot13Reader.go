package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rotate13(s byte) byte {
	b := rune(s)
	if (b >= 'a' && b <= 'm') || (b >= 'A' && b <= 'M') {
		s += 13
	}
	if (b >= 'n' && b <= 'z') || (b >= 'N' && b <= 'Z') {
		s -= 13
	}
	return s
}

func (rot13 rot13Reader) Read(data []byte) (readed int, err error) {
	readed, err = rot13.r.Read(data)
	for i := 0; i < readed; i++ {
		data[i] = rotate13(data[i])
	}
	return readed, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
