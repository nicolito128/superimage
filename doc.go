/*
Package superimage

Import it:

	import "github.com/nicolito128/superimage/v3"

Example program:

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

	    // Buffer for store the image data
	    buf := new(bytes.Buffer)
	    // Encode writes the image into the buffer
	    // gopher is ".png", so options can be nil
	    err = superimage.Encode(buf, img, nil)
	    if err != nil {
	        panic(err)
	    }

	    // Writing the cute gopher
	    os.WriteFile("gopher.png", buf.Bytes(), 0666)
	}
*/
package superimage
