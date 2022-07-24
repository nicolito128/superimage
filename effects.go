package superimage

import (
	"errors"
	"image"
	"image/color"
)

var ErrNegativeRadio = errors.New("Negative radio must be higher than 0.")
var ErrRadioTooBig = errors.New("Radio too big. It must be lower than the image's width or height.")

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

func Blur(img *SuperImage, radio int) (*SuperImage, error) {
	bounds := img.Bounds()
	width := img.Width
	height := img.Height

	if radio < 0 {
		return nil, ErrNegativeRadio
	}

	if radio >= width || radio >= height {
		return nil, ErrRadioTooBig
	}

	blurred := img.Factory(&image.NRGBA{}, nil).(*image.NRGBA)
	for x := bounds.Min.X + radio; x < width-radio; x++ {
		for y := bounds.Min.Y + radio; y < height-radio; y++ {
			i := blurred.PixOffset(x, y)
			p := blurred.Pix[i : i+4 : i+4]

			color := blurPointByRadio(img, x, y, radio)
			p[0] = color.R
			p[1] = color.G
			p[2] = color.B
			p[3] = color.A
		}
	}

	return New(blurred, img.Format), nil
}

func blurPointByRadio(img *SuperImage, x, y, radio int) color.NRGBA {
	if radio == 0 {
		r, g, b, a := img.At(x, y).RGBA()
		return color.NRGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}
	}

	var red, blue, green, alpha uint32

	for i := 1; i <= radio; i++ {
		r1, g1, b1, a1 := img.At(x, y+i).RGBA()
		r2, g2, b2, a2 := img.At(x, y-i).RGBA()
		r3, g3, b3, a3 := img.At(x-i, y).RGBA()
		r4, g4, b4, a4 := img.At(x+i, y).RGBA()
		r5, g5, b5, a5 := img.At(x+1, y+1).RGBA()
		r6, g6, b6, a6 := img.At(x-1, y-1).RGBA()
		r7, g7, b7, a7 := img.At(x, y).RGBA()

		red = (r1 + r2 + r3 + r4 + r5 + r6 + r7) / 7
		green = (g1 + g2 + g3 + g4 + g5 + g6 + g7) / 7
		blue = (b1 + b2 + b3 + b4 + b5 + b6 + b7) / 7
		alpha = (a1 + a2 + a3 + a4 + a5 + a6 + a7) / 7

	}

	return color.NRGBA{uint8(red >> 8), uint8(green >> 8), uint8(blue >> 8), uint8(alpha >> 8)}
}
