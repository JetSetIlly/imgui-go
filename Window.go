package imgui

// #include "wrapper/Window.h"
import "C"

// ShowDemoWindow creates a demo/test window. Demonstrates most ImGui features.
// Call this to learn about the library! Try to make it always available in your application!
func ShowDemoWindow(open *bool) {
	openArg, openFin := wrapBool(open)
	defer openFin()
	C.iggShowDemoWindow(openArg)
}

// ShowUserGuide adds basic help/info block (not a window): how to manipulate ImGui as a end-user (mouse/keyboard controls).
func ShowUserGuide() {
	C.iggShowUserGuide()
}

// WindowFlags for BeginV(), etc.
type WindowFlags int

const (
	WindowFlagsNone                      WindowFlags = 0
	WindowFlagsNoTitleBar                WindowFlags = 1 << 0  // Disable title-bar
	WindowFlagsNoResize                  WindowFlags = 1 << 1  // Disable user resizing with the lower-right grip
	WindowFlagsNoMove                    WindowFlags = 1 << 2  // Disable user moving the window
	WindowFlagsNoScrollbar               WindowFlags = 1 << 3  // Disable scrollbars (window can still scroll with mouse or programmatically)
	WindowFlagsNoScrollWithMouse         WindowFlags = 1 << 4  // Disable user vertically scrolling with mouse wheel. On child window, mouse wheel will be forwarded to the parent unless NoScrollbar is also set.
	WindowFlagsNoCollapse                WindowFlags = 1 << 5  // Disable user collapsing window by double-clicking on it. Also referred to as Window Menu Button (e.g. within a docking node).
	WindowFlagsAlwaysAutoResize          WindowFlags = 1 << 6  // Resize every window to its content every frame
	WindowFlagsNoBackground              WindowFlags = 1 << 7  // Disable drawing background color (WindowBg, etc.) and outside border. Similar as using SetNextWindowBgAlpha(0.0f).
	WindowFlagsNoSavedSettings           WindowFlags = 1 << 8  // Never load/save settings in .ini file
	WindowFlagsNoMouseInputs             WindowFlags = 1 << 9  // Disable catching mouse, hovering test with pass through.
	WindowFlagsMenuBar                   WindowFlags = 1 << 10 // Has a menu-bar
	WindowFlagsHorizontalScrollbar       WindowFlags = 1 << 11 // Allow horizontal scrollbar to appear (off by default). You may use SetNextWindowContentSize(ImVec2(width,0.0f)); prior to calling Begin() to specify width. Read code in imgui_demo in the "Horizontal Scrolling" section.
	WindowFlagsNoFocusOnAppearing        WindowFlags = 1 << 12 // Disable taking focus when transitioning from hidden to visible state
	WindowFlagsNoBringToFrontOnFocus     WindowFlags = 1 << 13 // Disable bringing window to front when taking focus (e.g. clicking on it or programmatically giving it focus)
	WindowFlagsAlwaysVerticalScrollbar   WindowFlags = 1 << 14 // Always show vertical scrollbar (even if ContentSize.y < Size.y)
	WindowFlagsAlwaysHorizontalScrollbar             = 1 << 15 // Always show horizontal scrollbar (even if ContentSize.x < Size.x)
	WindowFlagsNoNavInputs               WindowFlags = 1 << 16 // No keyboard/gamepad navigation within the window
	WindowFlagsNoNavFocus                WindowFlags = 1 << 17 // No focusing toward this window with keyboard/gamepad navigation (e.g. skipped by CTRL+TAB)
	WindowFlagsUnsavedDocument           WindowFlags = 1 << 18 // Display a dot next to the title. When used in a tab/docking context, tab is selected when clicking the X + closure is not assumed (will wait for user to stop submitting the tab). Otherwise closure is assumed when pressing the X, so if you keep submitting the tab may reappear at end of tab bar.
	WindowFlagsNoNav                     WindowFlags = WindowFlagsNoNavInputs | WindowFlagsNoNavFocus
	WindowFlagsNoDecoration              WindowFlags = WindowFlagsNoTitleBar | WindowFlagsNoResize | WindowFlagsNoScrollbar | WindowFlagsNoCollapse
	WindowFlagsNoInputs                  WindowFlags = WindowFlagsNoMouseInputs | WindowFlagsNoNavInputs | WindowFlagsNoNavFocus
)

// BeginV pushes a new window to the stack and start appending to it.
// You may append multiple times to the same window during the same frame.
// If the open argument is provided, the window can be closed, in which case the value will be false after the call.
//
// Returns false if the window is currently not visible.
// Regardless of the return value, End() must be called for each call to Begin().
func BeginV(id string, open *bool, flags WindowFlags) bool {
	idArg, idFin := wrapString(id)
	defer idFin()
	openArg, openFin := wrapBool(open)
	defer openFin()
	return C.iggBegin(idArg, openArg, C.int(flags)) != 0
}

// Begin calls BeginV(id, nil, 0).
func Begin(id string) bool {
	return BeginV(id, nil, 0)
}

// End closes the scope for the previously opened window.
// Every call to Begin() must be matched with a call to End().
func End() {
	C.iggEnd()
}

// ChildFlags for BeginChildV(), etc.
type ChildFlags int

const (
	ChildFlagsNone                   ChildFlags = 0
	ChildFlagsBorders                ChildFlags = 1 << 0 // Show an outer border and enable WindowPadding. (IMPORTANT: this is always == 1 == true for legacy reason)
	ChildFlagsAlwaysUseWindowPadding ChildFlags = 1 << 1 // Pad with style.WindowPadding even if no border are drawn (no padding by default for non-bordered child windows because it makes more sense)
	ChildFlagsResizeX                ChildFlags = 1 << 2 // Allow resize from right border (layout direction). Enable .ini saving (unless ImGuiWindowFlags_NoSavedSettings passed to window flags)
	ChildFlagsResizeY                ChildFlags = 1 << 3 // Allow resize from bottom border (layout direction). "
	ChildFlagsAutoResizeX            ChildFlags = 1 << 4 // Enable auto-resizing width. Read "IMPORTANT: Size measurement" details above.
	ChildFlagsAutoResizeY            ChildFlags = 1 << 5 // Enable auto-resizing height. Read "IMPORTANT: Size measurement" details above.
	ChildFlagsAlwaysAutoResize       ChildFlags = 1 << 6 // Combined with AutoResizeX/AutoResizeY. Always measure size even when child is hidden, always return true, always disable clipping optimization! NOT RECOMMENDED.
	ChildFlagsFrameStyle             ChildFlags = 1 << 7 // Style the child window like a framed item: use FrameBg, FrameRounding, FrameBorderSize, FramePadding instead of ChildBg, ChildRounding, ChildBorderSize, WindowPadding.
	ChildFlagsNavFlattened           ChildFlags = 1 << 8 // [BETA] Share focus scope, allow keyboard/gamepad navigation to cross over parent border to this child or between sibling child windows.
)

// BeginChildV pushes a new child to the stack and starts appending to it.
// flags are the WindowFlags to apply.
func BeginChildV(id string, size Vec2, border bool, flags ChildFlags) bool {
	idArg, idFin := wrapString(id)
	defer idFin()
	sizeArg, _ := size.wrapped()
	return C.iggBeginChild(idArg, sizeArg, castBool(border), C.int(flags)) != 0
}

// BeginChild calls BeginChildV(id, Vec2{0,0}, false, 0).
func BeginChild(id string) bool {
	return BeginChildV(id, Vec2{}, false, 0)
}

// EndChild closes the scope for the previously opened child.
// Every call to BeginChild() must be matched with a call to EndChild().
func EndChild() {
	C.iggEndChild()
}

// WindowPos returns the current window position in screen space.
// This is useful if you want to do your own drawing via the DrawList API.
func WindowPos() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggWindowPos(valueArg)
	valueFin()
	return value
}

// WindowSize returns the size of the current window.
func WindowSize() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggWindowSize(valueArg)
	valueFin()
	return value
}

// WindowWidth returns the width of the current window.
func WindowWidth() float32 {
	return float32(C.iggWindowWidth())
}

// WindowHeight returns the height of the current window.
func WindowHeight() float32 {
	return float32(C.iggWindowHeight())
}

// ContentRegionAvail returns the size of the content region that is available (based on the current cursor position).
func ContentRegionAvail() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggContentRegionAvail(valueArg)
	valueFin()
	return value
}

// SetNextWindowPosV sets next window position.
// Call before Begin(). Use pivot=(0.5,0.5) to center on given point, etc.
func SetNextWindowPosV(pos Vec2, cond Condition, pivot Vec2) {
	posArg, _ := pos.wrapped()
	pivotArg, _ := pivot.wrapped()
	C.iggSetNextWindowPos(posArg, C.int(cond), pivotArg)
}

// SetNextWindowPos calls SetNextWindowPosV(pos, 0, Vec{0,0}).
func SetNextWindowPos(pos Vec2) {
	SetNextWindowPosV(pos, 0, Vec2{})
}

// SetNextWindowCollapsed sets the next window collapsed state.
func SetNextWindowCollapsed(collapsed bool, cond Condition) {
	C.iggSetNextWindowCollapsed(castBool(collapsed), C.int(cond))
}

// SetNextWindowSizeV sets next window size.
// Set axis to 0.0 to force an auto-fit on this axis. Call before Begin().
func SetNextWindowSizeV(size Vec2, cond Condition) {
	sizeArg, _ := size.wrapped()
	C.iggSetNextWindowSize(sizeArg, C.int(cond))
}

// SetNextWindowSize calls SetNextWindowSizeV(size, 0).
func SetNextWindowSize(size Vec2) {
	SetNextWindowSizeV(size, 0)
}

// SetNextWindowSizeConstraints set next window size limits.
// Use -1,-1 on either X/Y axis to preserve the current size.
// Use callback to apply non-trivial programmatic constraints.
func SetNextWindowSizeConstraints(sizeMin Vec2, sizeMax Vec2) {
	sizeMinArg, _ := sizeMin.wrapped()
	sizeMaxArg, _ := sizeMax.wrapped()
	C.iggSetNextWindowSizeConstraints(sizeMinArg, sizeMaxArg)
}

// SetNextWindowContentSize sets next window content size (~ enforce the range of scrollbars).
// Does not include window decorations (title bar, menu bar, etc.).
// Set one axis to 0.0 to leave it automatic. This function must be called before Begin() to take effect.
func SetNextWindowContentSize(size Vec2) {
	sizeArg, _ := size.wrapped()
	C.iggSetNextWindowContentSize(sizeArg)
}

// SetNextWindowFocus sets next window to be focused / front-most. Call before Begin().
func SetNextWindowFocus() {
	C.iggSetNextWindowFocus()
}

// SetNextWindowBgAlpha sets next window background color alpha.
// Helper to easily modify ImGuiCol_WindowBg/ChildBg/PopupBg.
func SetNextWindowBgAlpha(value float32) {
	C.iggSetNextWindowBgAlpha(C.float(value))
}

// PushItemWidth pushes width of items for common large "item+label" widgets.
// >0.0f: width in pixels, <0.0f align xx pixels to the right of window (so -math.SmallestNonzeroFloat32 always align width to the right side).
func PushItemWidth(width float32) {
	C.iggPushItemWidth(C.float(width))
}

// PopItemWidth must be called for each call to PushItemWidth().
func PopItemWidth() {
	C.iggPopItemWidth()
}

// SetNextItemWidth sets width of the _next_ common large "item+label" widget.
// >0.0f: width in pixels, <0.0f align xx pixels to the right of window (so -math.SmallestNonzeroFloat32 always align width to the right side).
func SetNextItemWidth(width float32) {
	C.iggSetNextItemWidth(C.float(width))
}

// ItemFlags for PushItemFlag().
type ItemFlags int

const (
	ItemFlagsNone              ItemFlags = 0      // (Default)
	ItemFlagsNoTabStop         ItemFlags = 1 << 0 // false    // Disable keyboard tabbing. This is a "lighter" version of ItemFlagsNoNav.
	ItemFlagsNoNav             ItemFlags = 1 << 1 // false    // Disable any form of focusing (keyboard/gamepad directional navigation and SetKeyboardFocusHere() calls).
	ItemFlagsNoNavDefaultFocus ItemFlags = 1 << 2 // false    // Disable item being a candidate for default focus (e.g. used by title bar items).
	ItemFlagsButtonRepeat      ItemFlags = 1 << 3 // false    // Any button-like behavior will have repeat mode enabled (based on io.KeyRepeatDelay and io.KeyRepeatRate values). Note that you can also call IsItemActive() after any button to tell if it is being held.
	ItemFlagsAutoClosePopups   ItemFlags = 1 << 4 // true     // MenuItem()/Selectable() automatically close their parent popup window.
	ItemFlagsAllowDuplicateId  ItemFlags = 1 << 5 // false    // Allow submitting an item with the same identifier as an item already submitted this frame without triggering a warning tooltip if io.ConfigDebugHighlightIdConflicts is set.
)

// PushItemFlag changes flags in the existing options for the next items until PopItemFlag() is called.
func PushItemFlag(options ItemFlags, enabled bool) {
	C.iggPushItemFlag(C.int(options), castBool(enabled))
}

// PopItemFlag restores flags that were changed by the previous call to PushItemFlag().
func PopItemFlag() {
	C.iggPopItemFlag()
}

// CalcItemWidth returns the width of items given pushed settings and current cursor position.
func CalcItemWidth() float32 {
	return float32(C.iggCalcItemWidth())
}

// PushTextWrapPosV defines word-wrapping for Text() commands.
// < 0.0f: no wrapping; 0.0f: wrap to end of window (or column); > 0.0f: wrap at 'wrapPosX' position in window local space.
// Requires a matching call to PopTextWrapPos().
func PushTextWrapPosV(wrapPosX float32) {
	C.iggPushTextWrapPos(C.float(wrapPosX))
}

// PushTextWrapPos calls PushTextWrapPosV(0).
func PushTextWrapPos() {
	PushTextWrapPosV(0)
}

// PopTextWrapPos resets the last pushed position.
func PopTextWrapPos() {
	C.iggPopTextWrapPos()
}

// Viewport A Platform Window (always only one in 'master' branch), in the future may represent Platform Monitor.
type Viewport uintptr

// ViewportFlags flags for viewport.
type ViewportFlags int

const (
	ViewportFlagsNone              ViewportFlags = 0
	ViewportFlagsIsPlatformWindow  ViewportFlags = 1 << 0 // Represent a Platform Window
	ViewportFlagsIsPlatformMonitor ViewportFlags = 1 << 1 // Represent a Platform Monitor (unused yet)
	ViewportFlagsOwnedByApp        ViewportFlags = 1 << 2 // Platform Window: Is created/managed by the application (rather than a dear imgui backend)
)

// MainViewport returns primary/default viewport.
// - Currently represents the Platform Window created by the application which is hosting our Dear ImGui windows.
// - In 'docking' branch with multi-viewport enabled, we extend this concept to have multiple active viewports.
// - In the future we will extend this concept further to also represent Platform Monitor and support a "no main platform window" operation mode.
func MainViewport() Viewport {
	return Viewport(C.iggGetMainViewport())
}

func (viewport Viewport) handle() C.IggViewport {
	return C.IggViewport(viewport)
}

// Flags returns viewports flags value.
func (viewport Viewport) Flags() ViewportFlags {
	if viewport == 0 {
		return ViewportFlagsNone
	}
	return ViewportFlags(C.iggViewportGetFlags(viewport.handle()))
}

// Pos returns viewports Main Area: Position of the viewport (Dear Imgui coordinates are the same as OS desktop/native coordinates).
func (viewport Viewport) Pos() Vec2 {
	if viewport == 0 {
		return Vec2{}
	}

	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggViewportGetPos(viewport.handle(), valueArg)
	valueFin()
	return value
}

// Size returns viewports Main Area: Size of the viewport.
func (viewport Viewport) Size() Vec2 {
	if viewport == 0 {
		return Vec2{}
	}

	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggViewportGetSize(viewport.handle(), valueArg)
	valueFin()
	return value
}

// WorkPos returns viewports Work Area: Position of the viewport minus task bars, menus bars, status bars (>= Pos).
func (viewport Viewport) WorkPos() Vec2 {
	if viewport == 0 {
		return Vec2{}
	}

	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggViewportGetWorkPos(viewport.handle(), valueArg)
	valueFin()
	return value
}

// WorkSize returns viewports Work Area: Size of the viewport minus task bars, menu bars, status bars (<= Size).
func (viewport Viewport) WorkSize() Vec2 {
	if viewport == 0 {
		return Vec2{}
	}

	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggViewportGetWorkSize(viewport.handle(), valueArg)
	valueFin()
	return value
}

// Center returns center of the viewport.
func (viewport Viewport) Center() Vec2 {
	if viewport == 0 {
		return Vec2{}
	}

	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggViewportGetCenter(viewport.handle(), valueArg)
	valueFin()
	return value
}

// WorkCenter returns center of the viewport minus task bars, menu bars, status bars.
func (viewport Viewport) WorkCenter() Vec2 {
	if viewport == 0 {
		return Vec2{}
	}

	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggViewportGetWorkCenter(viewport.handle(), valueArg)
	valueFin()
	return value
}
