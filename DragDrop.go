package imgui

// #include "wrapper/DragDrop.h"
import "C"

// DragDropFlags for BeginDragDropSource(), etc.
type DragDropFlags int

const (
	DragDropFlagsNone DragDropFlags = 0
	// BeginDragDropSource() flags
	DragDropFlagsSourceNoPreviewTooltip   DragDropFlags = 1 << 0 // Disable preview tooltip. By default, a successful call to BeginDragDropSource opens a tooltip so you can display a preview or description of the source contents. This flag disables this behavior.
	DragDropFlagsSourceNoDisableHover     DragDropFlags = 1 << 1 // By default, when dragging we clear data so that IsItemHovered() will return false, to avoid subsequent user code submitting tooltips. This flag disables this behavior so you can still call IsItemHovered() on the source item.
	DragDropFlagsSourceNoHoldToOpenOthers DragDropFlags = 1 << 2 // Disable the behavior that allows to open tree nodes and collapsing header by holding over them while dragging a source item.
	DragDropFlagsSourceAllowNullID        DragDropFlags = 1 << 3 // Allow items such as Text(), Image() that have no unique identifier to be used as drag source, by manufacturing a temporary identifier based on their window-relative position. This is extremely unusual within the dear imgui ecosystem and so we made it explicit.
	DragDropFlagsSourceExtern             DragDropFlags = 1 << 4 // External source (from outside of dear imgui), won't attempt to read current item/window info. Will always return true. Only one Extern source can be active simultaneously.
	DragDropFlagsPayloadAutoExpire        DragDropFlags = 1 << 5 // Automatically expire the payload if the source cease to be submitted (otherwise payloads are persisting while being dragged)
	DragDropFlagsPayloadNoCrossContext    DragDropFlags = 1 << 6 // Hint to specify that the payload may not be copied outside current dear imgui context.
	DragDropFlagsPayloadNoCrossProcess    DragDropFlags = 1 << 7 // Hint to specify that the payload may not be copied outside current process.
	// AcceptDragDropPayload() flags
	DragDropFlagsAcceptBeforeDelivery    DragDropFlags = 1 << 10                                                                  // AcceptDragDropPayload() will returns true even before the mouse button is released. You can then call IsDelivery() to test if the payload needs to be delivered.
	DragDropFlagsAcceptNoDrawDefaultRect DragDropFlags = 1 << 11                                                                  // Do not draw the default highlight rectangle when hovering over target.
	DragDropFlagsAcceptNoPreviewTooltip  DragDropFlags = 1 << 12                                                                  // Request hiding the BeginDragDropSource tooltip from the BeginDragDropTarget site.
	DragDropFlagsAcceptPeekOnly          DragDropFlags = DragDropFlagsAcceptBeforeDelivery | DragDropFlagsAcceptNoDrawDefaultRect // For peeking ahead and inspecting the payload before delivery.
)

const (
	// DragDropPayloadTypeColor3F is payload type for 3 floats component color.
	DragDropPayloadTypeColor3F = "_COL3F"
	// DragDropPayloadTypeColor4F is payload type for 4 floats component color.
	DragDropPayloadTypeColor4F = "_COL4F"
)

// BeginDragDropSource registers the currently active item as drag'n'drop source.
// When this returns true you need to:
// a) call SetDragDropPayload() exactly once,
// b) you may render the payload visual/description,
// c) call EndDragDropSource().
func BeginDragDropSource(flags DragDropFlags) bool {
	return C.iggBeginDragDropSource(C.int(flags)) != 0
}

// SetDragDropPayload sets the payload for current draw and drop source.
// Strings starting with '_' are reserved for dear imgui internal types.
// Data is copied and held by imgui.
func SetDragDropPayload(dataType string, data []byte, cond Condition) bool {
	typeArg, typeFin := wrapString(dataType)
	defer typeFin()
	dataArg, dataFin := wrapBytes(data)
	defer dataFin()
	return C.iggSetDragDropPayload(typeArg, dataArg, C.int(len(data)), C.int(cond)) != 0
}

// EndDragDropSource closes the scope for current draw and drop source.
// Only call EndDragDropSource() if BeginDragDropSource() returns true.
func EndDragDropSource() {
	C.iggEndDragDropSource()
}

// BeginDragDropTarget must be called after submitting an item that may receive an item.
// If this returns true, you can call AcceptDragDropPayload() and EndDragDropTarget().
func BeginDragDropTarget() bool {
	return C.iggBeginDragDropTarget() != 0
}

// AcceptDragDropPayload accepts contents of a given type.
// If ImGuiDragDropFlags_AcceptBeforeDelivery is set you can peek into the payload before the mouse button is released.
func AcceptDragDropPayload(dataType string, flags DragDropFlags) []byte {
	typeArg, typeFin := wrapString(dataType)
	defer typeFin()

	payload := C.iggAcceptDragDropPayload(typeArg, C.int(flags))
	if payload == nil {
		return nil
	}

	data := C.iggPayloadData(payload)
	size := C.iggPayloadDataSize(payload)
	return C.GoBytes(data, size)
}

// EndDragDropTarget closed the scope for current drag and drop target.
// Only call EndDragDropTarget() if BeginDragDropTarget() returns true.
func EndDragDropTarget() {
	C.iggEndDragDropTarget()
}
