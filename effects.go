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

// Flip inverts the image horizontally returning a new *SuperImage.
func Flip(img *SuperImage) *SuperImage {
	bounds := img.Bounds()
	width := img.Width
	height := img.Height

	flipped := img.Factory(&image.RGBA{}, nil).(*image.RGBA)
	for x := bounds.Min.X; x < width; x++ {
		for y := bounds.Min.Y; y <= height/2; y++ {
			oppositeY := height - y - 1

			// Top quadrant index
			i := flipped.PixOffset(x, y)
			// Bottom quadrant index
			j := flipped.PixOffset(x, oppositeY)

			// Top quadrant pixel
			p := flipped.Pix[i : i+4]
			// Bottom quadrant pixel
			q := flipped.Pix[j : j+4]

			// Pixel colors
			r1, g1, b1, a1 := img.At(x, y).RGBA()
			r2, g2, b2, a2 := img.At(x, oppositeY).RGBA()

			// Parsing colors to uint8
			c1 := color.RGBA{uint8(r1 >> 8), uint8(g1 >> 8), uint8(b1 >> 8), uint8(a1 >> 8)}
			c2 := color.RGBA{uint8(r2 >> 8), uint8(g2 >> 8), uint8(b2 >> 8), uint8(a2 >> 8)}

			// Assigning colors to quadrants
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

// Reflect inverts the image vertically returning a new *SuperImage.
func Reflect(img *SuperImage) *SuperImage {
	bounds := img.Bounds()
	width := img.Width
	height := img.Height

	reflected := img.Factory(&image.RGBA{}, nil).(*image.RGBA)
	for y := bounds.Min.Y; y < height; y++ {
		for x := bounds.Min.X; x <= width/2; x++ {
			// X point of the reflected image
			oppositeX := width - x - 1

			// Left quadrant index
			i := reflected.PixOffset(x, y)
			// Right quadrant index
			j := reflected.PixOffset(oppositeX, y)

			// Left quadrant pixel
			p := reflected.Pix[i : i+4]
			// Right quadrant pixel
			q := reflected.Pix[j : j+4]

			// Pixel colors
			r1, g1, b1, a1 := img.At(x, y).RGBA()
			r2, g2, b2, a2 := img.At(oppositeX, y).RGBA()

			// Parsing colors to uint8
			c1 := color.RGBA{uint8(r1 >> 8), uint8(g1 >> 8), uint8(b1 >> 8), uint8(a1 >> 8)}
			c2 := color.RGBA{uint8(r2 >> 8), uint8(g2 >> 8), uint8(b2 >> 8), uint8(a2 >> 8)}

			// Assigning colors to quadrants
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

	return New(reflected, img.Format)
}
