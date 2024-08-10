package main

import (
	"bytes"
	"log"
	"os"
	"time"

	"github.com/nicolito128/superimage/v2"
)

func main() {
	log.Println("Starting opacity-gopher example...")
	start := time.Now()
	defer func() {
		log.Printf("Time since example started: %dms\n", time.Since(start).Milliseconds())
	}()

	img, err := superimage.GetByURL("https://go.dev/blog/gopher/gopher.png")
	if err != nil {
		panic(err)
	}

	// Buffer for store the image data
	buf := new(bytes.Buffer)
	// Encode writes the image into the buffer
	// gopher is ".png", so options can be nil
	err = superimage.Encode(buf, img, nil, nil)
	if err != nil {
		panic(err)
	}

	// Opacity returns a new *SuperImage with the image opacity set to the given value.
	// The opacity must be between 0 and 1, otherwise an error is returned.
	opimg, err := superimage.Opacity(img, 0.5)
	if err != nil {
		panic(err)
	}

	// Encoding on the buffer
	buf = new(bytes.Buffer)
	err = superimage.Encode(buf, opimg, nil, nil)
	if err != nil {
		panic(err)
	}

	// Writing the cute transparent gopher
	file, err := os.Create("examples/opacity/gopher.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Write(buf.Bytes())
}
