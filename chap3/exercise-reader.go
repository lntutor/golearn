package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (reader MyReader) Read(input []byte) (int, error) {
	for i := range input {
		input[i] = 65
	}
	return len(input), nil
}

func main() {
	reader.Validate(MyReader{})
}
