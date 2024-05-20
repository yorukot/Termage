package termage

import (
	"bytes"
	"image"
	"os"

	"github.com/BourgeoisBear/rasterm"
)

func showImageInKittyA(path string, maxWidth, maxHeight int, noresize bool) (imageString string, err error) {

	imageFile, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer imageFile.Close()

	img, _, err := image.Decode(imageFile)
	if err != nil {
		return "", err
	}

	img, err = resizeImage(img, maxWidth, maxHeight, noresize)
	if err != nil {
		return "", err
	}

	var buff bytes.Buffer
	err = rasterm.KittyWriteImage(&buff, img, rasterm.KittyImgOpts{})
	if err != nil {
		return "", err
	}

	return buff.String(), nil
}

func resizeTest(path string, maxWidth, maxHeight int, noresize bool) (imageString string, err error) {

	imageFile, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer imageFile.Close()

	img, _, err := image.Decode(imageFile)
	if err != nil {
		return "", err
	}

	_, err = resizeImage(img, maxWidth, maxHeight, noresize)
	if err != nil {
		return "", err
	}
	return "", err
}
