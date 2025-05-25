package superimage

import (
	"image"
	"image/color"
	"runtime"
	"sync"
)

func setPixelColors(p []uint8, c color.Color) {
	r, g, b, a := c.RGBA()
	p[0], p[1], p[2], p[3] = uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8)
}

func colorToUint8(c color.Color) color.Color {
	r, g, b, a := c.RGBA()
	return color.NRGBA{
		R: uint8(r >> 8),
		G: uint8(g >> 8),
		B: uint8(b >> 8),
		A: uint8(a >> 8),
	}
}

func transformColor(c color.Color, handler func(pix [4]uint32) [4]uint32) color.Color {
	r, g, b, a := c.RGBA()
	pix := handler([4]uint32{(r), (g), (b), (a)})
	return color.NRGBA{uint8(pix[0] >> 8), uint8(pix[1] >> 8), uint8(pix[2] >> 8), uint8(pix[3] >> 8)}
}

func getWorkers(limit int) (workers, linesPerWorker int) {
	workers = min(limit, runtime.NumCPU())
	linesPerWorker = limit / workers
	return
}

// Negative inverts the colors of an image.
func Negative(img image.Image) *SuperImage {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	inverted := image.NewNRGBA(bounds)

	var wg sync.WaitGroup
	numWorkers, linesPerWorker := getWorkers(height)

	for i := range numWorkers {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			startY := workerID * linesPerWorker
			endY := startY + linesPerWorker
			if workerID == numWorkers-1 {
				endY = height
			}

			for y := startY; y < endY; y++ {
				for x := range width {
					i := inverted.PixOffset(x, y)
					p := inverted.Pix[i : i+4 : i+4]

					c := transformColor(img.At(x, y), func(pix [4]uint32) [4]uint32 {
						return [4]uint32{
							pix[3] - pix[0],
							pix[3] - pix[1],
							pix[3] - pix[2],
							pix[3],
						}
					})

					setPixelColors(p, c)
				}
			}
		}(i)
	}
	wg.Wait()

	sp, ok := img.(SuperImage)
	if ok {
		return New(inverted, sp.Format())
	}

	return New(inverted, "png")
}

// Flip inverts the image horizontally returning a new *SuperImage.
func Flip(img image.Image) *SuperImage {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	flipped := image.NewRGBA(bounds)

	var wg sync.WaitGroup
	numWorkers, linesPerWorker := getWorkers(height)

	for i := range numWorkers {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			startY := workerID * linesPerWorker
			endY := startY + linesPerWorker
			if workerID == numWorkers-1 {
				endY = height
			}

			for y := startY; y < endY; y++ {
				for x := range width {
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
					pix0 := img.At(x, y)
					pix1 := img.At(x, oppositeY)

					// Parsing colors to uint8
					c1 := colorToUint8(pix0)
					c2 := colorToUint8(pix1)

					// Assigning colors to quadrants
					setPixelColors(p, c2)
					setPixelColors(q, c1)
				}
			}
		}(i)
	}
	wg.Wait()

	sp, ok := img.(SuperImage)
	if ok {
		return New(flipped, sp.Format())
	}

	return New(flipped, "png")
}

// Reflect inverts the image vertically returning a new *SuperImage.
func Reflect(img image.Image) *SuperImage {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	reflected := image.NewRGBA(bounds)

	var wg sync.WaitGroup
	numWorkers, linesPerWorker := getWorkers(height)

	for i := range numWorkers {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			startY := workerID * linesPerWorker
			endY := startY + linesPerWorker
			if workerID == numWorkers-1 {
				endY = height
			}

			for y := startY; y < endY; y++ {
				for x := range width {
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
					p[0], p[1], p[2], p[3] = c2.R, c2.G, c2.B, c2.A
					q[0], q[1], q[2], q[3] = c1.R, c1.G, c1.B, c1.A
				}
			}
		}(i)
	}
	wg.Wait()

	sp, ok := img.(SuperImage)
	if ok {
		return New(reflected, sp.Format())
	}

	return New(reflected, "png")
}

// Blur blurs an image by a given radio.
// If the radio is negative or bigger than the image's width or height, it returns an error.
// Radio 0 returns the original image without any change.
//
// References: https://relate.cs.illinois.edu/course/cs357-f15/file-version/03473f64afb954c74c02e8988f518de3eddf49a4/media/00-python-numpy/Image%20Blurring.html | http://arantxa.ii.uam.es/~jms/pfcsteleco/lecturas/20081215IreneBlasco.pdf
func Blur(img image.Image, radio int) (*SuperImage, error) {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	if radio < 0 {
		return nil, ErrNegativeRadio
	}
	blurred := image.NewNRGBA(img.Bounds())

	var wg sync.WaitGroup
	numWorkers, linesPerWorker := getWorkers(height)

	for i := range numWorkers {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			startY := workerID * linesPerWorker
			endY := startY + linesPerWorker
			if workerID == numWorkers-1 {
				endY = height
			}

			for y := startY; y < endY; y++ {
				for x := range width {
					i := blurred.PixOffset(x, y)
					p := blurred.Pix[i : i+4 : i+4]

					r, g, b, a := img.At(x, y).RGBA()
					p[0] = uint8(r >> 8)
					p[1] = uint8(g >> 8)
					p[2] = uint8(b >> 8)
					p[3] = uint8(a >> 8)
				}
			}
		}(i)
	}
	wg.Wait()

	for i := radio; i > 0; i-- {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for x := bounds.Min.X; x < width-1; x++ {
				for y := bounds.Min.Y; y < height-1; y++ {
					i := blurred.PixOffset(x, y)
					p := blurred.Pix[i : i+4 : i+4]

					r1, g1, b1, a1 := blurred.At(x, y).RGBA()
					r2, g2, b2, a2 := blurred.At(x-1, y).RGBA()
					r3, g3, b3, a3 := blurred.At(x+1, y).RGBA()
					r4, g4, b4, a4 := blurred.At(x, y-1).RGBA()
					r5, g5, b5, a5 := blurred.At(x, y+1).RGBA()

					p[0] = uint8(((r1*4 + r2 + r3 + r4 + r5) / 8) >> 8)
					p[1] = uint8(((g1*4 + g2 + g3 + g4 + g5) / 8) >> 8)
					p[2] = uint8(((b1*4 + b2 + b3 + b4 + b5) / 8) >> 8)
					p[3] = uint8(((a1*4 + a2 + a3 + a4 + a5) / 8) >> 8)
				}
			}
		}(i)
	}
	wg.Wait()

	sp, ok := img.(SuperImage)
	if ok {
		return New(blurred, sp.Format()), nil
	}

	return New(blurred, "png"), nil
}

func Opacity(img image.Image, op float64) (*SuperImage, error) {
	if op > 1 || op < 0 {
		return nil, ErrInvalidOpacity
	}

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	edited := image.NewNRGBA(bounds)

	var wg sync.WaitGroup
	numWorkers, linesPerWorker := getWorkers(height)

	for i := range numWorkers {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			startY := workerID * linesPerWorker
			endY := startY + linesPerWorker
			if workerID == numWorkers-1 {
				endY = height
			}

			for y := startY; y < endY; y++ {
				for x := range width {
					i := edited.PixOffset(x, y)
					p := edited.Pix[i : i+4 : i+4]

					r, g, b, a := img.At(x, y).RGBA()
					p[0] = uint8(r >> 8)
					p[1] = uint8(g >> 8)
					p[2] = uint8(b >> 8)
					p[3] = uint8(uint32(float64(a)*op) >> 8)
				}
			}
		}(i)
	}
	wg.Wait()

	sp, ok := img.(SuperImage)
	if ok {
		return New(edited, sp.Format()), nil
	}

	return New(edited, "png"), nil
}
