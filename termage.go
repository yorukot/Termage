package termage

import (
	"errors"
	"os"
	"strings"
)

type terminalType int

const (
    KittyProtocol terminalType = iota
    Unknown
)

func ImagePreview(path string, width, height int, noresize, fullwidth bool) (image string, err error) {
    currentTerminal := detectCurrentTerminal()
    if currentTerminal == Unknown {
        return "", errors.New("unknown/unsupported terminal")
    }
    
	if currentTerminal == KittyProtocol {
		return showImageInKitty(path, width, height, noresize, fullwidth)
    }

    return "", errors.New("unknown/unsupported terminal")
}


func detectCurrentTerminal() terminalType {
    term := os.Getenv("TERM")
    switch strings.ToLower(term) {
    case "xterm-kitty":
        return KittyProtocol
    case "konsole":
        return KittyProtocol
    case "xterm-256color":
        return KittyProtocol
    case "wezterm":
        return KittyProtocol
    case "mintty":
        return KittyProtocol
    case "foot":
        return KittyProtocol
    case "ghostty":
        return KittyProtocol
    case "blackbox":
        return KittyProtocol
    case "visualstudio":
        return KittyProtocol
    case "tabby":
        return KittyProtocol
    case "hyper":
        return KittyProtocol
    default:
        return Unknown
    }
}

