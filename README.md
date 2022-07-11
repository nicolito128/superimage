# SuperImage
Go module that provides some useful functions for working with images.

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
    - [Using `ParseURL`](#using-parseurl)
    - [Using `Negative`](#using-negative)
    - [Using `Flip`](#using-flip)
    - [Using `Reflect`](#using-reflect)
  - [Interest links](#interest-links)

## Getting Started

### Installation
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

### Quick start
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

## Examples
You have some good examples on how to use the package in the `examples/` folder.

## References

### About `SuperImage`
SuperImage is an Go struct, it can be used as any **Image** from the `std image package` because it's an Image interface implementation with the methods _At()_, _Bounds()_ and _ColorModel()_. You can create a new SuperImage with the _New(...)_ function.

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

    println(urlImg.Width, urlImg.Height)
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

    println(fileImg.Width, fileImg.Height)
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

	println(i.Height, i.Width)
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
	err = superimage.Encode(buf, img, nil)
	if err != nil {
		panic(err)
	}

    println(len(buf.Bytes()))
}
```

### Using `ParseURL`
```go
func main() {
    url, format, err := superimage.ParseURL("./project/images/cool.jpg")
    if err != nil {
        panic(err)
    }

    println(url.Path, format)
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
    err = superimage.Encode(buf, neg, nil)
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
    err = superimage.Encode(buf, flipped, nil)
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
    err = superimage.Encode(buf, reflect, nil)
    if err != nil {
        panic(err)
    }

    ioutil.WriteFile("./reflect.png", buf.Bytes(), 0666)
}
```

## Interest links
* [Go image standard library](https://pkg.go.dev/image)