package superimage

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type EncodeOptions struct {
	// Encode quality for Jpeg format: 1-100
	Quality int
}

// GetImageFile gets an image from a current project file.
func GetByFile(filename string) (*SuperImage, error) {
	// Getting the file format.
	_, format, err := ParseURL(filename)
	if err != nil {
		return nil, err
	}

	// Open the file with the right permissions.
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
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
	_, format, err := ParseURL(link)
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

	b, err := ioutil.ReadAll(res.Body)
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
func ParseURL(link string) (u *url.URL, format string, err error) {
	u, err = url.Parse(link)
	if err != nil {
		return u, "", err
	}

	format = u.Path[len(u.Path)-4:]
	if strings.Contains(format, ".") {
		format = strings.Split(format, ".")[1]
	}

	if format != "png" && format != "jpg" && format != "jpeg" {
		return u, "", fmt.Errorf("Unsupported format: %s", format)
	}

	return u, format, nil
}

// Decode decodes an image from r using the specified format (png, jpg, jpeg).
func Decode(r io.Reader, format string) (*SuperImage, error) {
	var img image.Image
	if format == "png" {
		im, err := png.Decode(r)
		img = im
		if err != nil {
			return nil, err
		}

	} else if format == "jpg" || format == "jpeg" {
		im, err := jpeg.Decode(r)
		img = im
		if err != nil {
			return nil, err
		}

	} else {
		return nil, fmt.Errorf("Unsupported format: %s", format)
	}

	return New(img, format), nil
}

// Encode writes the Image m to the given writer in the specified format (png, jpg, jpeg).
// If op *EncodeOptions is nil, the default options used are: { Format: png, Quality: 100 }.
func Encode(w io.Writer, m *SuperImage, op *EncodeOptions) error {
	if op == nil {
		op = &EncodeOptions{
			Quality: 100,
		}
	}

	switch m.Format {
	case "png":
		return png.Encode(w, m)
	case "jpg", "jpeg":
		return jpeg.Encode(w, m, &jpeg.Options{Quality: op.Quality})
	default:
		return fmt.Errorf("Unsupported format: %s", m.Format)
	}
}
