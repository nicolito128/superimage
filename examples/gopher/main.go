package main

import (
	"bytes"
	"log"
	"os"
	"time"

	"github.com/nicolito128/superimage/v3"
)

func main() {
	log.Println("Starting gopher example...")
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
	err = superimage.Encode(buf, img, nil)
	if err != nil {
		panic(err)
	}

	// Writing the cute gopher
	file, err := os.Create("examples/gopher/gopher.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Write(buf.Bytes())
}
