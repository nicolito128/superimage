# SuperImage
Go module that provides some useful functions for working with images.

## Installation
To install SuperImage package, you need solve the following issues:

1. Install [Go](https://go.dev/) (**version 1.18.+ required**).

2. Import it in your code:
```go
    import "github.com/nicolito128/superimage"
```

3. Get the package using go modules:
```
    go get github.com/nicolito/superimage
```

## Quick start
`project/main.go`:
```go
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

    // Writing the cute gopher
    ioutil.WriteFile("gopher.png", buf.Bytes(), 0666)
}
```