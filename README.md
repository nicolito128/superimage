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
      - [Using `GetByURL`](#using-getbyurl)
      - [Using `GetByFile`](#using-getbyfile)
    - [Using `Decode`](#using-decode)
    - [Using `Encode`](#using-encode)
    - [Using `ParseURL`](#using-parseurl)
    - [Using `Negative`](#using-negative)
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

#### Using `GetByURL`
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

#### Using `GetByFile`
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
```go
func main() {
    file, _ := os.Open("./examples/gopher/gopher.png")
	i, _ := superimage.Decode(file, "png")
	println(i.Height, i.Width)
}
```

### Using `Encode`

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

## Interest links
* [Go image standard library](https://pkg.go.dev/image)