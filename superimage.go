package superimage

import (
	"image"
)

// SuperImage is an image.Image implementation that wraps another image.Image.
type SuperImage struct {
	image.Image

	// Image format: png, jpg, jpeg.
	Format string
}

func New(im image.Image, format string) *SuperImage {
	return &SuperImage{im, format}
}
