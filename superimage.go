package superimage

import (
	"image"
)

// SuperImage is an image.Image implementation that wraps another image.Image.
type SuperImage struct {
	// Image format: png, jpg, jpeg.
	Format string
	image.Image
}

func New(im image.Image, format string) *SuperImage {
	return &SuperImage{format, im}
}
