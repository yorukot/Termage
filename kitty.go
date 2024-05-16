package termage

import (
	"encoding/base64"
	"fmt"
	"image"
	"os"
)

const (
	chunck_size = 4096
	start       = "\x1b_G"
	end         = "\x1b\\"
)

func showImageInKitty(path string, width, height int, noresize, fullwidth bool) (imageString string, err error) {
	width *= 10
	height *= 10

	imageFile, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer imageFile.Close()

	img, _, err := image.Decode(imageFile)
	if err != nil {
		return "", err
	}

	img, err = resizeImage(img, width, height, noresize, fullwidth)
	if err != nil {
		return "", err
	}

	imageFormat, err := getImageFormat(img)
	if err != nil {
		return "", err
	}

	imageBytes, err := imageToBytes(img, imageFormat)
	if err != nil {
		return "", err
	}
	
	imageBase64 := base64.StdEncoding.EncodeToString(imageBytes)

	imageString = fmt.Sprintf("%sa=T,f=100;%s%s", start, imageBase64, end)
	return imageString, nil
}
