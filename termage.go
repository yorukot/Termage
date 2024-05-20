package termage

import (
	"errors"
	"os"
	"strings"
)

type ImageProtocolType int

const (
	KittyProtocol ImageProtocolType = iota
	InlineImages
	Unknown
)

func ImagePreview(path string, maxWidth, maxHeight int, noresize bool) (image string, err error) {
	ImageProtocol := detectCurrentTerminalImageProtocol()
	if ImageProtocol == Unknown {
		return "", errors.New("unknown/unsupported terminal")
	}

	switch ImageProtocol {
		case KittyProtocol:
			return showImageInKittyA(path, maxWidth, maxHeight, noresize)
	}

	return "", errors.New("unknown/unsupported terminal")
}

func detectCurrentTerminalImageProtocol() ImageProtocolType {
	term := os.Getenv("TERM")
	switch strings.ToLower(term) {
	case "xterm-kitty":
		return KittyProtocol
	case "konsole":
		return KittyProtocol
	case "xterm-256color":
		return InlineImages
	case "wezterm":
		return InlineImages
	case "mintty":
		return InlineImages
	case "foot":
		return InlineImages
	case "ghostty":
		return KittyProtocol
	case "visualstudio":
		return InlineImages
	case "tabby":
		return InlineImages
	case "hyper":
		return InlineImages
	default:
		return Unknown
	}
}
