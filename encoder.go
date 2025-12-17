package superimage

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

var (
	DefaultPNGEncoder = &png.Encoder{
		CompressionLevel: png.NoCompression,
	}

	DefaultEncodeOptions = &EncodeOptions{
		PngEnc:   DefaultPNGEncoder,
		JpegOpts: nil,
		GifOpts:  nil,
	}
)

type EncodeOptions struct {
	PngEnc   *png.Encoder
	JpegOpts *jpeg.Options
	GifOpts  *gif.Options
}

// Encode writes the Image m to the given writer in the specified format (png, jpg/jpeg, gif).
func Encode(w io.Writer, m image.Image, opts *EncodeOptions) error {
	if opts == nil {
		opts = DefaultEncodeOptions
	}

	// Default format to encode
	format := "png"

	sp, ok := (m).(SuperImage)
	if ok {
		format = sp.Format()
	}

	switch format {
	case "png":
		return opts.PngEnc.Encode(w, m)
	case "jpg", "jpeg":
		return jpeg.Encode(w, m, opts.JpegOpts)
	case "gif":
		return gif.Encode(w, m, opts.GifOpts)
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}
}
