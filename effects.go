package superimage

import (
	"image"
	"image/color"
)

// Negative inverts the colors of an image returning a new image.Image interface.
func Negative(img *SuperImage) *SuperImage {
	bounds := img.Bounds()
	width := img.Width
	height := img.Height
	inverted := img.Factory(&image.NRGBA{}, nil).(*image.NRGBA)

	for y := bounds.Min.Y; y < height; y++ {
		for x := bounds.Min.X; x < width; x++ {
			i := inverted.PixOffset(x, y)
			p := inverted.Pix[i : i+4 : i+4]

			r, g, b, a := img.At(x, y).RGBA()
			c := color.NRGBA{uint8((a - r) >> 8), uint8((a - g) >> 8), uint8((a - b) >> 8), uint8(a >> 8)}

			p[0] = c.R
			p[1] = c.G
			p[2] = c.B
			p[3] = c.A
		}
	}

	return New(inverted, img.Format)
}
