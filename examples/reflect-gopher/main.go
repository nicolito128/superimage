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

	// Buffer store the image data
	buf := new(bytes.Buffer)
	// Encode writes the image into the buffer
	// gopher is ".png", so options can be nil
	err = superimage.Encode(buf, img, nil)
	if err != nil {
		panic(err)
	}

	// Reflect the image vertically
	reflect := superimage.Reflect(img)
	// Encoding on the buffer
	buf = new(bytes.Buffer)
	err = superimage.Encode(buf, reflect, nil)
	if err != nil {
		panic(err)
	}

	// Writing the cute reflected gopher
	ioutil.WriteFile("examples/reflect-gopher/gopher.png", buf.Bytes(), 0666)
}
