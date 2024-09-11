package termage

import (
	"bytes"
	"image"
	"os"

	"github.com/BourgeoisBear/rasterm"
)

func showImageInKittyA(path string, maxWidth, maxHeight int, noresize bool) (string, error) {
	
	img, err := loadImage(path)
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

	imageFile, err := loadImage(path)
	if err != nil {
		return "", err
	}

	//change _ to resizedImg in case you want to get back the img
	_, err = resizeImage(imageFile, maxWidth, maxHeight, noresize)
	if err != nil {
		return "", err
	}

	// In case you might want to do something with the resized image
	// For testing purposes we'll just return its dimensions as a string
	// return fmt.Sprintf("Resized image dimensions: %dx%d", resizedImg.Bounds().Dx(), resizedImg.Bounds().Dy()), nil
}

func loadImage(path string) (image.Image, error) {
	imageFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer imageFile.Close()

	img, _, err := image.Decode(imageFile)
	return img, err
}
