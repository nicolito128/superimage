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

	// Blur returns a new *SuperImage with the image blurred by radio.
	// Higher radio means more blur.
	// If radio is 0, the image is not blurred.
	blurred, err := superimage.Blur(img, 2)
	if err != nil {
		panic(err)
	}

	// Encoding on the buffer
	buf = new(bytes.Buffer)
	err = superimage.Encode(buf, blurred, nil)
	if err != nil {
		panic(err)
	}

	// Writing the cute blurred gopher
	ioutil.WriteFile("examples/blur-gopher/gopher.png", buf.Bytes(), 0666)
}
