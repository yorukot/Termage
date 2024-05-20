package termage

import (
	"fmt"
	"testing"
)

func TestAddImagePreview(t *testing.T) {
	imageString, err := ImagePreview("./testFile/1580624931717m.jpg", 10, 10, false)
	if err != nil {
		t.Errorf("Failed to generate image preview: %v", err)
		return
	}
	fmt.Println(imageString)
}

func TestAddResize(t *testing.T) {
	_, err := resizeTest("./testFile/1580624931717m.jpg", 70, 50, false)
	if err != nil {
		t.Errorf("Failed to resize image: %v", err)
		return
	}
}
