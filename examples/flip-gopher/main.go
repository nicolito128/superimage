package main

import (
	"bytes"
	"os"

	"github.com/nicolito128/superimage/v2"
)

func main() {
	img, err := superimage.GetByURL("https://go.dev/blog/gopher/gopher.png")
	if err != nil {
		panic(err)
	}

	// Buffer store the image data
	buf := new(bytes.Buffer)
	// Encode writes the image into the buffer
	// gopher is ".png", so options can be nil
	err = superimage.Encode(buf, img, nil, nil)
	if err != nil {
		panic(err)
	}

	// Flip the image horizontally
	flip := superimage.Flip(img)
	// Encoding on the buffer
	buf = new(bytes.Buffer)
	err = superimage.Encode(buf, flip, nil, nil)
	if err != nil {
		panic(err)
	}

	// Writing the cute flipped gopher
	file, err := os.Create("examples/flip-gopher/gopher.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Write(buf.Bytes())
}
