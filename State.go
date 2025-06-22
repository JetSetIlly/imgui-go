package imgui

// #include "wrapper/State.h"
import "C"

// ClearActiveID removes focus from a currently item in edit mode.
// An application that handles its own undo/redo stack needs to call this
// function before changing the data a widget might currently own, such as
// a TextEdit().
func ClearActiveID() {
	C.iggClearActiveID()
}

// IsItemClicked returns true if the current item is clicked with the left
// mouse button.
func IsItemClicked() bool {
	return C.iggIsItemClicked() != 0
}

// IsItemHoveredV returns true if the last item is hovered.
// (and usable, aka not blocked by a popup, etc.). See HoveredFlags for more options.
func IsItemHoveredV(flags HoveredFlags) bool {
	return C.iggIsItemHovered(C.int(flags)) != 0
}

// IsItemHovered calls IsItemHoveredV(HoveredFlagsNone).
func IsItemHovered() bool {
	return IsItemHoveredV(HoveredFlagsNone)
}

// IsItemActive returns if the last item is active.
// e.g. button being held, text field being edited.
//
// This will continuously return true while holding mouse button on an item.
// Items that don't interact will always return false.
func IsItemActive() bool {
	return C.iggIsItemActive() != 0
}

// IsItemEdited return true if the last item was modified or was pressed. This
// is generally the same as the "bool" return value of many widgets.
func IsItemEdited() bool {
	return C.iggIsItemEdited() != 0
}

// IsItemActivated returns true if the last item was made active (item was
// previously inactive).
func IsItemActivated() bool {
	return C.iggIsItemActivated() != 0
}

// IsItemDeactivated returns true if the last item was made inactive (item was
// previously active).
func IsItemDeactivated() bool {
	return C.iggIsItemDeactivated() != 0
}

// IsItemDeactivatedAfterEdit returns true if the last item was made inactive
// and made a value change when it was active (e.g. Slider/Drag moved).
func IsItemDeactivatedAfterEdit() bool {
	return C.iggIsItemDeactivatedAfterEdit() != 0
}

// IsItemToggledOpen returns true if the last item's open was toggled open.
func IsItemToggledOpen() bool {
	return C.iggIsItemToggledOpen() != 0
}

// IsAnyItemActive returns true if the any item is active.
func IsAnyItemActive() bool {
	return C.iggIsAnyItemActive() != 0
}

// IsItemVisible returns true if the last item is visible.
func IsItemVisible() bool {
	return C.iggIsItemVisible() != 0
}

// SetItemAllowOverlap allows last item to be overlapped by a subsequent item.
// Both may be activated during the same frame before the later one takes priority.
// This is sometimes useful with invisible buttons, selectables, etc. to catch unused area.
// [deprecated]
func SetItemAllowOverlap() {
	C.iggSetNextItemAllowOverlap()
}

// SetItemAllowOverlap allows last item to be overlapped by a subsequent item.
// Both may be activated during the same frame before the later one takes priority.
// This is sometimes useful with invisible buttons, selectables, etc. to catch unused area.
func SetNextItemAllowOverlap() {
	C.iggSetNextItemAllowOverlap()
}

// IsWindowAppearing returns whether the current window is appearing.
func IsWindowAppearing() bool {
	return C.iggIsWindowAppearing() != 0
}

// IsWindowCollapsed returns whether the current window is collapsed.
func IsWindowCollapsed() bool {
	return C.iggIsWindowCollapsed() != 0
}

// FocusedFlags for IsWindowFocusedV().
type FocusedFlags int

// This is a list of FocusedFlags combinations.
const (
	FocusedFlagsNone             FocusedFlags = 0
	FocusedFlagsChildWindows     FocusedFlags = 1 << 0 // Return true if any children of the window is focused
	FocusedFlagsRootWindow       FocusedFlags = 1 << 1 // Test from root window (top most parent of the current hierarchy)
	FocusedFlagsAnyWindow        FocusedFlags = 1 << 2 // Return true if any window is focused. Important: If you are trying to tell how to dispatch your low-level inputs, do NOT use this. Use 'io.WantCaptureMouse' instead! Please read the FAQ!
	FocusedFlagsNoPopupHierarchy FocusedFlags = 1 << 3 // Do not consider popup hierarchy (do not treat popup emitter as parent of popup) (when used with _ChildWindows or _RootWindow)
	//FocusedFlagsDockHierarchy               FocusedFlags = 1 << 4   // Consider docking hierarchy (treat dockspace host as parent of docked window) (when used with _ChildWindows or _RootWindow)
	FocusedFlagsRootAndChildWindows FocusedFlags = FocusedFlagsRootWindow | FocusedFlagsChildWindows
)

// IsWindowFocusedV returns if current window is focused or its root/child, depending on flags. See flags for options.
func IsWindowFocusedV(flags FocusedFlags) bool {
	return C.iggIsWindowFocused(C.int(flags)) != 0
}

// IsWindowFocused calls IsWindowFocusedV(FocusedFlagsNone).
func IsWindowFocused() bool {
	return IsWindowFocusedV(FocusedFlagsNone)
}

// HoveredFlags for IsWindowHoveredV(), etc.
type HoveredFlags int

// This is a list of HoveredFlags combinations.
const (
	HoveredFlagsNone             HoveredFlags = 0      // Return true if directly over the item/window, not obstructed by another window, not obstructed by an active popup or modal blocking inputs under them.
	HoveredFlagsChildWindows     HoveredFlags = 1 << 0 // IsWindowHovered() only: Return true if any children of the window is hovered
	HoveredFlagsRootWindow       HoveredFlags = 1 << 1 // IsWindowHovered() only: Test from root window (top most parent of the current hierarchy)
	HoveredFlagsAnyWindow        HoveredFlags = 1 << 2 // IsWindowHovered() only: Return true if any window is hovered
	HoveredFlagsNoPopupHierarchy HoveredFlags = 1 << 3 // IsWindowHovered() only: Do not consider popup hierarchy (do not treat popup emitter as parent of popup) (when used with _ChildWindows or _RootWindow)
	//HoveredFlagsDockHierarchy               HoveredFlags = 1 << 4   // IsWindowHovered() only: Consider docking hierarchy (treat dockspace host as parent of docked window) (when used with _ChildWindows or _RootWindow)
	HoveredFlagsAllowWhenBlockedByPopup HoveredFlags = 1 << 5 // Return true even if a popup window is normally blocking access to this item/window
	//HoveredFlagsAllowWhenBlockedByModal     HoveredFlags = 1 << 6   // Return true even if a modal popup window is normally blocking access to this item/window. FIXME-TODO: Unavailable yet.
	HoveredFlagsAllowWhenBlockedByActiveItem HoveredFlags = 1 << 7  // Return true even if an active item is blocking access to this item/window. Useful for Drag and Drop patterns.
	HoveredFlagsAllowWhenOverlappedByItem    HoveredFlags = 1 << 8  // IsItemHovered() only: Return true even if the item uses AllowOverlap mode and is overlapped by another hoverable item.
	HoveredFlagsAllowWhenOverlappedByWindow  HoveredFlags = 1 << 9  // IsItemHovered() only: Return true even if the position is obstructed or overlapped by another window.
	HoveredFlagsAllowWhenDisabled            HoveredFlags = 1 << 10 // IsItemHovered() only: Return true even if the item is disabled
	HoveredFlagsNoNavOverride                HoveredFlags = 1 << 11 // IsItemHovered() only: Disable using keyboard/gamepad navigation state when active, always query mouse
	HoveredFlagsAllowWhenOverlapped          HoveredFlags = HoveredFlagsAllowWhenOverlappedByItem | HoveredFlagsAllowWhenOverlappedByWindow
	HoveredFlagsRectOnly                     HoveredFlags = HoveredFlagsAllowWhenBlockedByPopup | HoveredFlagsAllowWhenBlockedByActiveItem | HoveredFlagsAllowWhenOverlapped
	HoveredFlagsRootAndChildWindows          HoveredFlags = HoveredFlagsRootWindow | HoveredFlagsChildWindows

	// Tooltips mode
	// - typically used in IsItemHovered() + SetTooltip() sequence.
	// - this is a shortcut to pull flags from 'style.HoverFlagsForTooltipMouse' or 'style.HoverFlagsForTooltipNav' where you can reconfigure desired behavior.
	//   e.g. 'TooltipHoveredFlagsForMouse' defaults to 'HoveredFlagsStationary | HoveredFlagsDelayShort'.
	// - for frequently actioned or hovered items providing a tooltip you want may to use HoveredFlagsForTooltip (stationary + delay) so the tooltip doesn't show too often.
	// - for items which main purpose is to be hovered or items with low affordance, or in less consistent apps, prefer no delay or shorter delay.
	HoveredFlagsForTooltip HoveredFlags = 1 << 12 // Shortcut for standard flags when using IsItemHovered() + SetTooltip() sequence.

	// (Advanced) Mouse Hovering delays.
	// - generally you can use HoveredFlagsForTooltip to use application-standardized flags.
	// - use those if you need specific overrides.
	HoveredFlagsStationary    HoveredFlags = 1 << 13 // Require mouse to be stationary for style.HoverStationaryDelay (~0.15 sec) _at least one time_. After this, can move on same item/window. Using the stationary test tends to reduces the need for a long delay.
	HoveredFlagsDelayNone     HoveredFlags = 1 << 14 // IsItemHovered() only: Return true immediately (default). As this is the default you generally ignore this.
	HoveredFlagsDelayShort    HoveredFlags = 1 << 15 // IsItemHovered() only: Return true after style.HoverDelayShort elapsed (~0.15 sec) (shared between items) + requires mouse to be stationary for style.HoverStationaryDelay (once per item).
	HoveredFlagsDelayNormal   HoveredFlags = 1 << 16 // IsItemHovered() only: Return true after style.HoverDelayNormal elapsed (~0.40 sec) (shared between items) + requires mouse to be stationary for style.HoverStationaryDelay (once per item).
	HoveredFlagsNoSharedDelay HoveredFlags = 1 << 17 // IsItemHovered() only: Disable shared delay system where moving from one item to the next keeps the previous timer for a short time (standard for tooltips with long delays)
)

// IsWindowHoveredV returns if current window is hovered (and typically: not blocked by a popup/modal).
// See flags for options. NB: If you are trying to check whether your mouse should be dispatched to imgui or to your app,
// you should use the 'io.WantCaptureMouse' boolean for that!
func IsWindowHoveredV(flags HoveredFlags) bool {
	return C.iggIsWindowHovered(C.int(flags)) != 0
}

// IsWindowHovered calls IsWindowHoveredV(HoveredFlagsNone).
func IsWindowHovered() bool {
	return IsWindowHoveredV(HoveredFlagsNone)
}

// IsKeyDown returns true if the corresponding key is currently being held down.
func IsKeyDown(key int) bool {
	return C.iggIsKeyDown(C.int(key)) != 0
}

// IsKeyPressedV returns true if the corresponding key was pressed (went from !Down to Down).
// If repeat=true and the key is being held down then the press is repeated using io.KeyRepeatDelay and KeyRepeatRate.
func IsKeyPressedV(key int, repeat bool) bool {
	return C.iggIsKeyPressed(C.int(key), castBool(repeat)) != 0
}

// IsKeyPressed calls IsKeyPressedV(key, true).
func IsKeyPressed(key int) bool {
	return IsKeyPressedV(key, true)
}

// IsKeyReleased returns true if the corresponding key was released (went from Down to !Down).
func IsKeyReleased(key int) bool {
	return C.iggIsKeyReleased(C.int(key)) != 0
}

// IsMouseDown returns true if the corresponding mouse button is currently being held down.
func IsMouseDown(button int) bool {
	return C.iggIsMouseDown(C.int(button)) != 0
}

// IsAnyMouseDown returns true if any mouse button is currently being held down.
func IsAnyMouseDown() bool {
	return C.iggIsAnyMouseDown() != 0
}

// IsMouseClickedV returns true if the mouse button was clicked (0=left, 1=right, 2=middle)
// If repeat=true and the mouse button is being held down then the click is repeated using io.KeyRepeatDelay and KeyRepeatRate.
func IsMouseClickedV(button int, repeat bool) bool {
	return C.iggIsMouseClicked(C.int(button), castBool(repeat)) != 0
}

// IsMouseClicked calls IsMouseClickedV(key, false).
func IsMouseClicked(button int) bool {
	return IsMouseClickedV(button, false)
}

// IsMouseReleased returns true if the mouse button was released (went from Down to !Down).
func IsMouseReleased(button int) bool {
	return C.iggIsMouseReleased(C.int(button)) != 0
}

// IsMouseDoubleClicked returns true if the mouse button was double-clicked (0=left, 1=right, 2=middle).
func IsMouseDoubleClicked(button int) bool {
	return C.iggIsMouseDoubleClicked(C.int(button)) != 0
}

// IsMouseDragging returns true if the mouse button is being dragged.
func IsMouseDragging(button int, threshold float64) bool {
	return C.iggIsMouseDragging(C.int(button), C.float(threshold)) != 0
}

// MouseDragDelta returns the delta from the initial clicking position while the mouse button is pressed or was just released. This is locked and return 0.0f until the mouse moves past a distance threshold at least once (if lockThreshold < -1.0f, uses io.MouseDraggingThreshold).
func MouseDragDelta(button int, lockThreshold float32) Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggGetMouseDragDelta(valueArg, C.int(button), C.float(lockThreshold))
	valueFin()
	return value
}

// ResetMouseDragDelta resets the drag delta.
func ResetMouseDragDelta(button int) {
	C.iggResetMouseDragDelta(C.int(button))
}

// MousePos returns the current mouse position in screen space.
func MousePos() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggMousePos(valueArg)
	valueFin()
	return value
}

// MouseCursorID for SetMouseCursor().
//
// User code may request backend to display given cursor by calling SetMouseCursor(),
// which is why we have some cursors that are marked unused here.
type MouseCursorID int

const (
	// MouseCursorNone no mouse cursor.
	MouseCursorNone MouseCursorID = -1
	// MouseCursorArrow standard arrow mouse cursor.
	MouseCursorArrow MouseCursorID = 0
	// MouseCursorTextInput when hovering over InputText, etc.
	MouseCursorTextInput MouseCursorID = 1
	// MouseCursorResizeAll (Unused by imgui functions).
	MouseCursorResizeAll MouseCursorID = 2
	// MouseCursorResizeNS when hovering over an horizontal border.
	MouseCursorResizeNS MouseCursorID = 3
	// MouseCursorResizeEW when hovering over a vertical border or a column.
	MouseCursorResizeEW MouseCursorID = 4
	// MouseCursorResizeNESW when hovering over the bottom-left corner of a window.
	MouseCursorResizeNESW MouseCursorID = 5
	// MouseCursorResizeNWSE when hovering over the bottom-right corner of a window.
	MouseCursorResizeNWSE MouseCursorID = 6
	// MouseCursorHand (Unused by imgui functions. Use for e.g. hyperlinks).
	MouseCursorHand MouseCursorID = 7
	// MouseCursorCount is the number of defined mouse cursors.
	MouseCursorCount MouseCursorID = 8
)

// MouseCursor returns desired cursor type, reset in imgui.NewFrame(), this is updated during the frame.
// Valid before Render(). If you use software rendering by setting io.MouseDrawCursor ImGui will render those for you.
func MouseCursor() MouseCursorID {
	return MouseCursorID(C.iggGetMouseCursor())
}

// SetMouseCursor sets desired cursor type.
func SetMouseCursor(cursor MouseCursorID) {
	C.iggSetMouseCursor(C.int(cursor))
}

// ItemRectMax returns the lower-right bounding rectangle of the last item in screen space.
func ItemRectMax() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggGetItemRectMax(valueArg)
	valueFin()
	return value
}

// ItemRectMin returns the upper-left bounding rectangle of the last item in screen space.
func ItemRectMin() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggGetItemRectMin(valueArg)
	valueFin()
	return value
}
