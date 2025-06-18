package imgui

// #include "wrapper/ListClipper.h"
import "C"
import (
	"runtime/cgo"
)

// ListClipper contains the final results after ListClipperAll() has finished iterating.
type ListClipper struct {
	DisplayStart int
	DisplayEnd   int
	ItemsCount   int
	ItemsHeight  float32
}

//export listClipperDraw
func listClipperDraw(hook C.uintptr_t, i int) {
	h := cgo.Handle(hook)
	h.Value().(func(int))(i)
}

// ListClipperAll is used instead of Dear Imgui's Begin(), Step() and End() functions. The draw
// function is called every step of the ListClipper.
func ListClipperAll(itemsCount int, draw func(int)) ListClipper {
	drawHandler := cgo.NewHandle(draw)
	defer drawHandler.Delete()

	var results C.IggListClipperResults
	C.iggListClipperAll(&results, C.int(itemsCount), C.float(-1), C.uintptr_t(drawHandler))

	return ListClipper{
		DisplayStart: int(results.DisplayStart),
		DisplayEnd:   int(results.DisplayEnd),
		ItemsCount:   itemsCount,
		ItemsHeight:  float32(results.ItemsHeight),
	}
}
