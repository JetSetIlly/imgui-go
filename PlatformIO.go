package imgui

// #include "wrapper/PlatformIO.h"
import "C"

// PlatformIO is the platform (SLD, GLFW, etc.) specific IO for Imgui. Access via CurrentPlatformIO().
type PlatformIO struct {
	handle C.IggPlatformIO
}

// CurrentPlatformIO returns the current PlatformIO.
func CurrentPlatformIO() PlatformIO {
	return PlatformIO{handle: C.iggGetCurrentPlatformIO()}
}

// Clipboard describes the access to the text clipboard of the window manager.
type Clipboard interface {
	// Text returns the current text from the clipboard, if available.
	ClipboardText() (string, error)
	// SetText sets the text as the current text on the clipboard.
	SetClipboardText(value string)
}

var clipboard Clipboard
var freeClipboardString func()

// SetClipboard registers a clipboard for text copy/paste actions.
// If no clipboard is set, then a fallback implementation may be used, if available for the OS.
// To disable clipboard handling overall, pass nil as the Clipboard.
//
// Since ImGui queries the clipboard text via a return value, the wrapper has to hold the
// current clipboard text as a copy in memory. This memory will be freed at the next clipboard operation.
func (io PlatformIO) SetClipboard(board Clipboard) {
	if freeClipboardString != nil {
		freeClipboardString()
		freeClipboardString = nil
	}
	if board != nil {
		clipboard = board
		C.iggPlatformIoRegisterClipboardFunctions(io.handle)
	} else {
		clipboard = nil
		C.iggPlatformIoClearClipboardFunctions(io.handle)
	}
}

//export iggPlatformIoGetClipboardText
func iggPlatformIoGetClipboardText(_ C.IggPlatformIO) *C.char {
	if freeClipboardString != nil {
		freeClipboardString()
		freeClipboardString = nil
	}
	if clipboard == nil {
		return nil
	}
	text, err := clipboard.ClipboardText()
	if err != nil {
		return nil
	}
	textPtr, textFin := wrapString(text)
	freeClipboardString = textFin
	return textPtr
}

//export iggPlatformIoSetClipboardText
func iggPlatformIoSetClipboardText(_ C.IggPlatformIO, text *C.char) {
	if freeClipboardString != nil {
		freeClipboardString()
		freeClipboardString = nil
	}
	if clipboard == nil {
		return
	}
	clipboard.SetClipboardText(C.GoString(text))
}
