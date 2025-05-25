# SuperImage
The package provides some useful structures and functions for working with images. Apply effects such as blur or negative color to an image with a well-performing tool.

## Index
- [SuperImage](#superimage)
  - [Index](#index)
  - [Getting Started](#getting-started)
    - [Installation](#installation)
    - [Quick start](#quick-start)
  - [Examples](#examples)
  - [References](#references)
    - [About `SuperImage`](#about-superimage)
    - [Using `GetByURL`](#using-getbyurl)
    - [Using `GetByFile`](#using-getbyfile)
    - [Using `Decode`](#using-decode)
    - [Using `Encode`](#using-encode)
    - [Using `Negative`](#using-negative)
    - [Using `Flip`](#using-flip)
    - [Using `Reflect`](#using-reflect)
    - [Using `Blur`](#using-blur)
    - [Using `Pixelate`](#using-pixelate)
  - [Interest links](#interest-links)

## Getting Started

### Installation
To install SuperImage package, you need solve the following issues:

1. Install [Go](https://go.dev/) (**version 1.24.+ recommended**).

2. Get the package using go modules:
```
    go get -u github.com/nicolito/superimage/v2
```

3. Import it in your code:
```go
    import "github.com/nicolito128/superimage/v2"
```

### Quick start
`project/main.go`:
```go
package main

import (
    "bytes"
	"os"

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
    err = superimage.Encode(buf, img, nil, nil)
    if err != nil {
        panic(err)
    }

    // Writing the cute gopher
    os.WriteFile("gopher.png", buf.Bytes(), 0666)
}
```

## Examples
You have some good examples on how to use the package in the `examples/` folder.

## References

### About `SuperImage`
SuperImage is an Go struct, it can be used as any **Image** from the `std image package` because it's an Image composition. You can create a new SuperImage with the _New(...)_ function.

```go
func main() {
    rec := image.Rectangle{image.Point{0, 0}, image.Point{500, 500}}
    boringImg := image.NewRGBA(rec)
    superImg := superimage.New(boringImg, "png")

    println(superImg.Bounds())
}
```

### Using `GetByURL`
Get a new SuperImage with an URL.

```go
func main() {
    // Getting a new SuperImage with a link
    urlImg, err := superimage.GetByURL("https://awesomeurl.com/image.png")
    if err != nil {
        panic(err)
    }

    println(urlImg.Bounds())
}
```

### Using `GetByFile`
Get a new SuperImage with a project file image.

```go
func main() {
    // Getting a new SuperImage with a file
    fileImg, err := superimage.GetByFile("./folder/cool_image.jpg")
    if err != nil {
        panic(err)
    }

    println(fileImg.Bounds())
}
```

### Using `Decode`
Decodes an reader on a new SuperImage.

```go
func main() {
    file, _ := os.Open("./examples/gopher/gopher.png")
	i, err := superimage.Decode(file, "png")
    if err != nil {
        panic(err)
    }

	println(i.Bounds())
}
```

### Using `Encode`
Encodes a writer on a new SuperImage.

```go
func main() {
    fileImg, err := superimage.GetByFile("./folder/cool_image.jpg")
    if err != nil {
        panic(err)
    }

    buf := new(bytes.Buffer)
	err = superimage.Encode(buf, img, nil, nil)
	if err != nil {
		panic(err)
	}

    println(len(buf.Bytes()))
}
```

### Using `Negative`
Inverts the colors of an image.

```go
func main() {
    img, err := superimage.GetByURL("https://awesomeurl.com/image.png")
    if err != nil {
        panic(err)
    }

    // Inverting image colors
    neg := superimage.Negative(img)

    // Saving
    buf := new(bytes.Buffer)
    err = superimage.Encode(buf, neg, nil, nil)
    if err != nil {
        panic(err)
    }

    ioutil.WriteFile("./negative.png", buf.Bytes(), 0666)
}
```

### Using `Flip`
Turn an image upside down.

```go
func main() {
    img, err := superimage.GetByURL("https://awesomeurl.com/image.png")
    if err != nil {
        panic(err)
    }

    // Flipping image
    flipped := superimage.Flip(img)

    // Saving
    buf := new(bytes.Buffer)
    err = superimage.Encode(buf, flipped, nil, nil)
    if err != nil {
        panic(err)
    }

    ioutil.WriteFile("./flipped.png", buf.Bytes(), 0666)
}
```

### Using `Reflect`
Reflects an image vertically.

```go
func main() {
    img, err := superimage.GetByURL("https://awesomeurl.com/image.png")
    if err != nil {
        panic(err)
    }

    // Reflecting image
    reflect := superimage.Reflect(img)

    // Saving
    buf := new(bytes.Buffer)
    err = superimage.Encode(buf, reflect, nil, nil)
    if err != nil {
        panic(err)
    }

    ioutil.WriteFile("./reflect.png", buf.Bytes(), 0666)
}
```

### Using `Blur`
Blur an image by a given radio.
```go
func main() {
    img, err := superimage.GetByURL("https://awesomeurl.com/image.png")
    if err != nil {
        panic(err)
    }

    // Blur
    blurred, err := superimage.Blur(img, 2)
    if err != nil {
        panic(err)
    }

    // Saving
    buf := new(bytes.Buffer)
    err = superimage.Encode(buf, blurred, nil, nil)
    if err != nil {
        panic(err)
    }

    ioutil.WriteFile("./blurred.png", buf.Bytes(), 0666)
}
```

### Using `Pixelate`
Pixelate an image by a given radio.
```go
func main() {
    img, err := superimage.GetByURL("https://awesomeurl.com/image.png")
    if err != nil {
        panic(err)
    }

    // Pixelate
    pixelated, err := superimage.Pixelate(img, 2)
    if err != nil {
        panic(err)
    }

    // Saving
    buf := new(bytes.Buffer)
    err = superimage.Encode(buf, pixelated, nil, nil)
    if err != nil {
        panic(err)
    }

    ioutil.WriteFile("./pixelated.png", buf.Bytes(), 0666)
}
```

## Interest links
* [Go image standard library](https://pkg.go.dev/image)