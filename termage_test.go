package termage

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	imageString, err := ImagePreview("./testFile/computer.png", 100, 100, false, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(imageString)
}