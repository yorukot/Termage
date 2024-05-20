package termage

import (
	"image"
	"math"

	"github.com/nfnt/resize"
)

func resizeImage(img image.Image, maxWidth, maxHeight int, noresize bool) (image.Image, error) {
	maxWidth *= 10
	maxHeight *= 10
	if noresize {
		return img, nil
	}

	currentWidth := img.Bounds().Dx()
	currentHeight := img.Bounds().Dy()

	if maxWidth > 0 && maxHeight > 0 && (currentWidth > maxWidth || currentHeight > maxHeight) {
		var scale float64
		if currentWidth > maxWidth || currentHeight > maxHeight {
			scale = math.Min(float64(maxWidth)/float64(currentWidth), float64(maxHeight)/float64(currentHeight))
		} else {
			scale = 1.0
		}

		newWidth := int(float64(currentWidth) * scale)
		newHeight := int(float64(currentHeight) * scale)

		resizedImg := resize.Resize(uint(newWidth), uint(newHeight), img, resize.Lanczos3)
		return resizedImg, nil
	}

	return img, nil
}
