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

// GetImageFile gets an image from a current project file.
func GetByFile(filename string) (*SuperImage, error) {
	// Getting the file format.
	_, format, err := parseURL(filename)
	if err != nil {
		return nil, err
	}

	// Open the file with the right permissions.
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

// GetImageFromURL gets an image from an URL with an http GET request.
func GetByURL(link string) (*SuperImage, error) {
	_, format, err := parseURL(link)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return nil, err
	}

	// Required to make a request.
	req.Close = true
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

// ParseURL parses an URL.
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

// Decode decodes an image from r using the specified format (png, jpg, jpeg).
func Decode(r io.Reader, format string) (*SuperImage, error) {
	var img image.Image

	switch format {
	case "png":
		im, err := png.Decode(r)
		img = im
		if err != nil {
			return nil, err
		}

	case "jpg", "jpeg":
		im, err := jpeg.Decode(r)
		img = im
		if err != nil {
			return nil, err
		}

	case "gif":
		im, err := gif.Decode(r)
		img = im
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}

	return New(img, format), nil
}

// Encode writes the Image m to the given writer in the specified format (png, jpg/jpeg, gif).
// If your image isn't a jpeg or gif just pass nil.
func Encode(w io.Writer, m image.Image, jpgOptions *jpeg.Options, gifOptions *gif.Options) error {
	var format = "png"

	sp, ok := m.(SuperImage)
	if ok {
		format = sp.Format
	}

	switch format {
	case "png":
		return png.Encode(w, m)
	case "jpg", "jpeg":
		return jpeg.Encode(w, m, jpgOptions)
	case "gif":
		return gif.Encode(w, m, gifOptions)
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}
}
