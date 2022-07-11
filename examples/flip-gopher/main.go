package main

import (
	"bytes"
	"io/ioutil"

	"github.com/nicolito128/superimage"
)

func main() {
	img, err := superimage.GetByURL("https://go.dev/blog/gopher/gopher.png")
	if err != nil {
		panic(err)
	}

	// Buffer for store the image data
	buf := new(bytes.Buffer)
	// Encode writes the image into the buffer
	// gopher is ".png", so options can be nil
	err = superimage.Encode(buf, img, nil)
	if err != nil {
		panic(err)
	}

	// Flip the image vertically
	neg := superimage.Flip(img)
	// Encoding on the buffer
	buf = new(bytes.Buffer)
	err = superimage.Encode(buf, neg, nil)
	if err != nil {
		panic(err)
	}

	// Writing the cute negative gopher
	ioutil.WriteFile("examples/flip-gopher/gopher.png", buf.Bytes(), 0666)
}
