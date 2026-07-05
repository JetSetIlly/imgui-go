package imgui

// #include "wrapper/Dock.h"
import "C"

func DockSpaceOverViewport() {
	C.iggDockSpaceOverViewport()
}
