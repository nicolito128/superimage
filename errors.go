package superimage

import "errors"

var (
	ErrNegativeRadio  = errors.New("radio must be higher than 0")
	ErrInvalidOpacity = errors.New("opacity must be between 0 and 1")
)
