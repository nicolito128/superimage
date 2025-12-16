package superimage

import (
	"image"
	"image/color"
	"runtime"
	"sync"
)

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
					j := inverted.PixOffset(x, y)
					p := inverted.Pix[j : j+4 : j+4]

					c := img.At(x, y)
					r, g, b, a := c.RGBA()

					invertedR := 0xFFFF - r
					invertedG := 0xFFFF - g
					invertedB := 0xFFFF - b

					p[0] = uint8(invertedR >> 8)
					p[1] = uint8(invertedG >> 8)
					p[2] = uint8(invertedB >> 8)
					p[3] = uint8(a >> 8)
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
					originalColor := img.At(x, y)
					flipped.Set(x, height-y-1, originalColor)
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
					originalColor := img.At(x, y)
					reflected.Set(width-x-1, y, originalColor)
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
func Blur(img image.Image, radius int) (*SuperImage, error) {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	if radius < 0 {
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

			// Ajusta startY/endY a los límites reales de la imagen
			if startY < bounds.Min.Y {
				startY = bounds.Min.Y
			}
			if endY > bounds.Max.Y {
				endY = bounds.Max.Y
			}

			for y := startY; y < endY; y++ {
				for x := bounds.Min.X; x < bounds.Max.X; x++ {
					var rSum, gSum, bSum, aSum uint64
					count := 0

					for ky := -radius; ky <= radius; ky++ {
						for kx := -radius; kx <= radius; kx++ {
							sampleX, sampleY := x+kx, y+ky

							if sampleX >= bounds.Min.X && sampleX < bounds.Max.X &&
								sampleY >= bounds.Min.Y && sampleY < bounds.Max.Y {
								c := img.At(sampleX, sampleY)
								r, g, b, a := c.RGBA()

								rSum += uint64(r)
								gSum += uint64(g)
								bSum += uint64(b)
								aSum += uint64(a)
								count++
							}
						}
					}

					var avgR, avgG, avgB, avgA uint32
					if count > 0 {
						avgR = uint32(rSum / uint64(count))
						avgG = uint32(gSum / uint64(count))
						avgB = uint32(bSum / uint64(count))
						avgA = uint32(aSum / uint64(count))
					}

					offset := blurred.PixOffset(x, y)
					blurred.Pix[offset+0] = uint8(avgR >> 8)
					blurred.Pix[offset+1] = uint8(avgG >> 8)
					blurred.Pix[offset+2] = uint8(avgB >> 8)
					blurred.Pix[offset+3] = uint8(avgA >> 8)
				}
			}
		}(i)
	}
	wg.Wait()

	for i := radius; i > 0; i-- {
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

	op16bit := uint32(op * 0xFFFF)

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
					_, _, _, curAlpha := img.At(x, y).RGBA()
					newAlpha := uint32((uint64(curAlpha) * uint64(op16bit)) / 0xFFFF)

					r, g, b, _ := img.At(x, y).RGBA()

					offset := edited.PixOffset(x, y)
					edited.Pix[offset+0] = uint8(r >> 8)
					edited.Pix[offset+1] = uint8(g >> 8)
					edited.Pix[offset+2] = uint8(b >> 8)
					edited.Pix[offset+3] = uint8(newAlpha >> 8)
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

func Pixelate(img image.Image, radius int) (*SuperImage, error) {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	pixelated := image.NewRGBA(bounds)

	var wg sync.WaitGroup

	numWorkers := runtime.NumCPU()
	numBlocksY := (height + radius - 1) / radius

	for i := range numWorkers {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			startBlockY := workerID * (numBlocksY / numWorkers)
			endBlockY := (workerID + 1) * (numBlocksY / numWorkers)
			if workerID == numWorkers-1 {
				endBlockY = numBlocksY
			}

			startY := startBlockY * radius
			endY := endBlockY * radius
			endY = min(endY, height)

			for y := startY; y < endY; y += radius {
				for x := 0; x < width; x += radius {
					blockRect := image.Rect(x, y, x+radius, y+radius)
					blockRect = blockRect.Intersect(bounds)

					avgColor := calculateAverageColourWithRect(img, blockRect)

					for dy := range radius {
						for dx := range radius {
							if x+dx < width && y+dy < height {
								pixelated.Set(x+dx, y+dy, avgColor)
							}
						}
					}
				}
			}
		}(i)
	}
	wg.Wait()

	sp, ok := img.(*SuperImage)
	if ok {
		return New(pixelated, sp.Format()), nil
	}

	return New(pixelated, "png"), nil
}

func calculateAverageColourWithRect(img image.Image, rect image.Rectangle) color.Color {
	var r, g, b, a uint32
	var count uint32

	imgBounds := img.Bounds()
	actualRect := rect.Intersect(imgBounds)

	for y := actualRect.Min.Y; y < actualRect.Max.Y; y++ {
		for x := actualRect.Min.X; x < actualRect.Max.X; x++ {
			c := img.At(x, y)
			rr, gg, bb, aa := c.RGBA()
			r += rr
			g += gg
			b += bb
			a += aa
			count++
		}
	}

	if count == 0 {
		return color.RGBA{0, 0, 0, 0}
	}

	return color.RGBA{
		R: uint8(r / (count) >> 8),
		G: uint8(g / (count) >> 8),
		B: uint8(b / (count) >> 8),
		A: uint8(a / (count) >> 8),
	}
}
