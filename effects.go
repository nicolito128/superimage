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

// Flip inverts the image vertically returning a new *SuperImage.
func Flip(img *SuperImage) *SuperImage {
	bounds := img.Bounds()
	width := img.Width
	height := img.Height

	flipped := img.Factory(&image.RGBA{}, nil).(*image.RGBA)
	for x := bounds.Min.X; x < width; x++ {
		for y := bounds.Min.Y; y < height/2; y++ {
			i := flipped.PixOffset(x, y)
			j := flipped.PixOffset(x, height-y-1)

			p := flipped.Pix[i : i+4]
			q := flipped.Pix[j : j+4]
			r1, g1, b1, a1 := img.At(x, y).RGBA()
			r2, g2, b2, a2 := img.At(x, height-y-1).RGBA()

			c1 := color.RGBA{uint8(r1 >> 8), uint8(g1 >> 8), uint8(b1 >> 8), uint8(a1 >> 8)}
			c2 := color.RGBA{uint8(r2 >> 8), uint8(g2 >> 8), uint8(b2 >> 8), uint8(a2 >> 8)}

			p[0] = c2.R
			p[1] = c2.G
			p[2] = c2.B
			p[3] = c2.A

			q[0] = c1.R
			q[1] = c1.G
			q[2] = c1.B
			q[3] = c1.A
		}
	}

	return New(flipped, img.Format)
}
