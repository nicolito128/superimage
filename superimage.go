package superimage

import (
	"image"
)

// SuperImage is an image.Image implementation
type SuperImage struct {
	image.Image

	// Image format: png, jpg, jpeg.
	format string
}

func New(im image.Image, format string) *SuperImage {
	return &SuperImage{
		Image:  im,
		format: format,
	}
}

func (si SuperImage) Format() string {
	return si.format
}
