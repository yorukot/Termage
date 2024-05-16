package termage

import (
	"bytes"
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/chai2010/webp"
	"golang.org/x/image/draw"
	"golang.org/x/term"
)

func resizeImage(img image.Image, width, height int, noresize, fullwidth bool) (image.Image, error) {
	if width <= 0 && height <= 0 && !fullwidth {
		return nil, errors.New("either width, height or fullwidth should be specified")
	}

	termWidth, _, err := terminalSize()
	if err != nil {
		return img, err
	}

	if !noresize {
		// Calculate new dimensions while maintaining aspect ratio
		if width > 0 && height == 0 {
			height = int(float64(img.Bounds().Dy()) * float64(width) / float64(img.Bounds().Dx()))
		} else if height > 0 && width == 0 {
			width = int(float64(img.Bounds().Dx()) * float64(height) / float64(img.Bounds().Dy()))
		} else if fullwidth {
			width = termWidth
			height = int(float64(img.Bounds().Dy()) * float64(termWidth) / float64(img.Bounds().Dx()))
		}
	}

	// Resize the image
	resizedImg := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.CatmullRom.Scale(resizedImg, resizedImg.Bounds(), img, img.Bounds(), draw.Over, nil)

	return resizedImg, nil
}

func terminalSize() (width, height int, err error) {
	width, height, err = term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 0, 0, err
	}
	return width, height, nil
}

func imageToBytes(img image.Image, format string) ([]byte, error) {
	var buffer bytes.Buffer

	// 选择合适的编码器
	switch format {
	case "jpeg":
		err := jpeg.Encode(&buffer, img, nil)
		if err != nil {
			return nil, err
		}
	case "png":
		err := png.Encode(&buffer, img)
		if err != nil {
			return nil, err
		}
	case "gif":
		err := gif.Encode(&buffer, img, nil)
		if err != nil {
			return nil, err
		}
	case "webp":
		options := &webp.Options{
			Lossless: false,
			Quality:  80,
		}
		err := webp.Encode(&buffer, img, options)
		if err != nil {
			return nil, err
		}

	default:
		return nil, errors.New("unsupported image format")
	}

	return buffer.Bytes(), nil
}

func getImageFormat(img image.Image) (string, error) {
    switch img.(type) {
    case *image.RGBA:
        return "png", nil
    case *image.YCbCr:
        return "jpeg", nil
    default:
        return "", errors.New("unknown image format")
    }
}