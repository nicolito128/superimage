package superimage

import (
	"image"
	"image/color"
)

// SuperImage is an image.Image implementation that wraps another image.Image.
type SuperImage struct {
	// self is the wrapped image.Image.
	self image.Image
	// Sizes of the image.
	Width  int
	Height int
	// Image format: png, jpg, jpeg.
	Format string
}

type FactoryOptions struct {
	YCbCrSubsampleRatio *image.YCbCrSubsampleRatio
	Palette             *color.Palette
	Uniform             *color.Color
}

func New(im image.Image, format string) *SuperImage {
	return &SuperImage{
		self:   im,
		Width:  im.Bounds().Dx(),
		Height: im.Bounds().Dy(),
		Format: format,
	}
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (s *SuperImage) At(x, y int) color.Color {
	return s.self.At(x, y)
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (s *SuperImage) Bounds() image.Rectangle {
	return s.self.Bounds()
}

// ColorModel returns the Image's color model.
func (s *SuperImage) ColorModel() color.Model {
	return s.self.ColorModel()
}

// Factory creates a new image of the same type as the model passed as argument.
// The function receives a FactoryOptions struct as argument, it can be nil. Also,
// any field of the struct can be nil, if the field is nil the function sets an
// default value.
//
// The FactoryOptions has the following fields by default:
//
// { YCbCrSubsampleRatio: 4, Palette: new(color.Pallete), Uniform: new(color.Color) }
func (s *SuperImage) Factory(model image.Image, op *FactoryOptions) image.Image {
	// Default options values.
	i := new(image.YCbCrSubsampleRatio)
	*i = 4
	p := new(color.Palette)
	u := new(color.Color)

	if op == nil {
		// Assignation
		op = &FactoryOptions{}
		op.YCbCrSubsampleRatio = i
		op.Palette = p
		op.Uniform = u
	} else {
		// FactoryOptions is not nil but some fields are nil.

		if op.YCbCrSubsampleRatio == nil {
			op.YCbCrSubsampleRatio = i
		}

		if op.Palette == nil {
			op.Palette = p
		}

		if op.Uniform == nil {
			op.Uniform = u
		}
	}

	// Create a new image.Image implementation.
	// with the model type and the options.
	switch model.(type) {
	case *image.RGBA:
		return image.NewRGBA(s.Bounds())
	case *image.NRGBA:
		return image.NewNRGBA(s.Bounds())
	case *image.RGBA64:
		return image.NewRGBA64(s.Bounds())
	case *image.NRGBA64:
		return image.NewNRGBA64(s.Bounds())
	case *image.Alpha:
		return image.NewAlpha(s.Bounds())
	case *image.Alpha16:
		return image.NewAlpha16(s.Bounds())
	case *image.Gray:
		return image.NewGray(s.Bounds())
	case *image.Gray16:
		return image.NewGray16(s.Bounds())
	case *image.CMYK:
		return image.NewCMYK(s.Bounds())
	case *image.YCbCr:
		return image.NewYCbCr(s.Bounds(), image.YCbCrSubsampleRatio(*op.YCbCrSubsampleRatio))
	case *image.NYCbCrA:
		return image.NewNYCbCrA(s.Bounds(), image.YCbCrSubsampleRatio(*op.YCbCrSubsampleRatio))
	case *image.Paletted:
		return image.NewPaletted(s.Bounds(), *op.Palette)
	case *image.Uniform:
		return image.NewUniform(*op.Uniform)
	default:
		return nil
	}
}
