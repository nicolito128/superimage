package superimage

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// GetByFile gets an image from a current project file.
func GetByFile(filename string) (*SuperImage, error) {
	// Getting the file format.
	_, format, err := parseURL(filename)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, err := Decode(file, format)
	if err != nil {
		return nil, err
	}

	return New(img, format), nil
}

// GetByURL gets an image from an URL with an http GET request.
func GetByURL(link string) (*SuperImage, error) {
	_, format, err := parseURL(link)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "image/"+format)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	img, err := Decode(bytes.NewReader(b), format)
	if err != nil {
		return nil, err
	}

	return New(img, format), nil
}

// Decode decodes an image from r using the specified format (png, jpg, jpeg, gif).
func Decode(r io.Reader, format string) (*SuperImage, error) {
	var img image.Image
	var err error

	switch format {
	case "png":
		img, err = png.Decode(r)
		if err != nil {
			return nil, err
		}

	case "jpg", "jpeg":
		img, err = jpeg.Decode(r)
		if err != nil {
			return nil, err
		}

	case "gif":
		img, err = gif.Decode(r)
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}

	return New(img, format), nil
}

// parseURL calls to the Parse method of the url package.
func parseURL(link string) (u *url.URL, format string, err error) {
	u, err = url.Parse(link)
	if err != nil {
		return u, "", err
	}

	format = u.Path[len(u.Path)-4:]
	if strings.Contains(format, ".") {
		format = strings.Split(format, ".")[1]
	}

	if format != "png" && format != "jpg" && format != "jpeg" {
		return u, "", fmt.Errorf("unsupported format: %s", format)
	}

	return u, format, nil
}
