package imgui

// #include "wrapper/Color.h"
import "C"

// ColorEditFlags for ColorEdit3V(), etc.
type ColorEditFlags int

const (
	ColorEditFlagsNone           ColorEditFlags = 0
	ColorEditFlagsNoAlpha        ColorEditFlags = 1 << 1  //              // ColorEdit, ColorPicker, ColorButton: ignore Alpha component (will only read 3 components from the input pointer).
	ColorEditFlagsNoPicker       ColorEditFlags = 1 << 2  //              // ColorEdit: disable picker when clicking on color square.
	ColorEditFlagsNoOptions      ColorEditFlags = 1 << 3  //              // ColorEdit: disable toggling options menu when right-clicking on inputs/small preview.
	ColorEditFlagsNoSmallPreview ColorEditFlags = 1 << 4  //              // ColorEdit, ColorPicker: disable color square preview next to the inputs. (e.g. to show only the inputs)
	ColorEditFlagsNoInputs       ColorEditFlags = 1 << 5  //              // ColorEdit, ColorPicker: disable inputs sliders/text widgets (e.g. to show only the small preview color square).
	ColorEditFlagsNoTooltip      ColorEditFlags = 1 << 6  //              // ColorEdit, ColorPicker, ColorButton: disable tooltip when hovering the preview.
	ColorEditFlagsNoLabel        ColorEditFlags = 1 << 7  //              // ColorEdit, ColorPicker: disable display of inline text label (the label is still forwarded to the tooltip and picker).
	ColorEditFlagsNoSidePreview  ColorEditFlags = 1 << 8  //              // ColorPicker: disable bigger color preview on right side of the picker, use small color square preview instead.
	ColorEditFlagsNoDragDrop     ColorEditFlags = 1 << 9  //              // ColorEdit: disable drag and drop target. ColorButton: disable drag and drop source.
	ColorEditFlagsNoBorder       ColorEditFlags = 1 << 10 //              // ColorButton: disable border (which is enforced by default)

	// Alpha preview
	// - Prior to 1.91.8 (2025/01/21): alpha was made opaque in the preview by default using old name ColorEditFlagsAlphaPreview.
	// - We now display the preview as transparent by default. You can use ColorEditFlagsAlphaOpaque to use old behavior.
	// - The new flags may be combined better and allow finer controls.
	ColorEditFlagsAlphaOpaque                    ColorEditFlags = 1 << 11 //              // ColorEdit, ColorPicker, ColorButton: disable alpha in the preview,. Contrary to _NoAlpha it may still be edited when calling ColorEdit4()/ColorPicker4(). For ColorButton() this does the same as _NoAlpha.
	ColorEditFlagsAlphaNoBg                      ColorEditFlags = 1 << 12 //              // ColorEdit, ColorPicker, ColorButton: disable rendering a checkerboard background behind transparent color.
	ColorEditFlagsAlphaPreviewHalfColorEditFlags                = 1 << 13 //              // ColorEdit, ColorPicker, ColorButton: display half opaque / half transparent preview.

	// User Options (right-click on widget to change some of them).
	ColorEditFlagsAlphaBar       ColorEditFlags = 1 << 16 //              // ColorEdit, ColorPicker: show vertical alpha bar/gradient in picker.
	ColorEditFlagsHDR            ColorEditFlags = 1 << 19 //              // (WIP) ColorEdit: Currently only disable 0.0f..1.0f limits in RGBA edition (note: you probably want to use ColorEditFlagsFloat flag as well).
	ColorEditFlagsDisplayRGB     ColorEditFlags = 1 << 20 // [Display]    // ColorEdit: override _display_ type among RGB/HSV/Hex. ColorPicker: select any combination using one or more of RGB/HSV/Hex.
	ColorEditFlagsDisplayHSV     ColorEditFlags = 1 << 21 // [Display]    // "
	ColorEditFlagsDisplayHex     ColorEditFlags = 1 << 22 // [Display]    // "
	ColorEditFlagsUint8          ColorEditFlags = 1 << 23 // [DataType]   // ColorEdit, ColorPicker, ColorButton: _display_ values formatted as 0..255.
	ColorEditFlagsFloat          ColorEditFlags = 1 << 24 // [DataType]   // ColorEdit, ColorPicker, ColorButton: _display_ values formatted as 0.0f..1.0f floats instead of 0..255 integers. No round-trip of value via integers.
	ColorEditFlagsPickerHueBar   ColorEditFlags = 1 << 25 // [Picker]     // ColorPicker: bar for Hue, rectangle for Sat/Value.
	ColorEditFlagsPickerHueWheel ColorEditFlags = 1 << 26 // [Picker]     // ColorPicker: wheel for Hue, triangle for Sat/Value.
	ColorEditFlagsInputRGB       ColorEditFlags = 1 << 27 // [Input]      // ColorEdit, ColorPicker: input and output data in RGB format.
	ColorEditFlagsInputHSV       ColorEditFlags = 1 << 28 // [Input]      // ColorEdit, ColorPicker: input and output data in HSV format.

	// Defaults Options. You can set application defaults using SetColorEditOptions(). The intent is that you probably don't want to
	// override them in most of your calls. Let the user choose via the option menu and/or call SetColorEditOptions() once during startup.
	ColorEditFlagsDefaultOptions_ = ColorEditFlagsUint8 | ColorEditFlagsDisplayRGB | ColorEditFlagsInputRGB | ColorEditFlagsPickerHueBar
)

// ColorEdit3 calls ColorEdit3V(label, col, 0).
func ColorEdit3(label string, col *[3]float32) bool {
	return ColorEdit3V(label, col, 0)
}

// ColorEdit3V will show a clickable little square which will open a color picker window for 3D vector (rgb format).
func ColorEdit3V(label string, col *[3]float32, flags ColorEditFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	ccol := (*C.float)(&col[0])
	return C.iggColorEdit3(labelArg, ccol, C.int(flags)) != 0
}

// ColorEdit4 calls ColorEdit4V(label, col, 0).
func ColorEdit4(label string, col *[4]float32) bool {
	return ColorEdit4V(label, col, 0)
}

// ColorEdit4V will show a clickable little square which will open a color picker window for 4D vector (rgba format).
func ColorEdit4V(label string, col *[4]float32, flags ColorEditFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	ccol := (*C.float)(&col[0])
	return C.iggColorEdit4(labelArg, ccol, C.int(flags)) != 0
}

// ColorButton displays a color square/button, hover for details, returns true when pressed.
func ColorButton(id string, col Vec4, flags ColorEditFlags, size Vec2) bool {
	idArg, idFin := wrapString(id)
	defer idFin()
	sizeArg, _ := size.wrapped()
	colArg, _ := col.wrapped()
	return C.iggColorButton(idArg, colArg, C.int(flags), sizeArg) != 0
}

// ColorPickerFlags for ColorPicker3V(), etc.
type ColorPickerFlags int

const (
	// ColorPickerFlagsNone default = 0.
	ColorPickerFlagsNone ColorPickerFlags = 0
	// ColorPickerFlagsNoPicker disables picker when clicking on colored square.
	ColorPickerFlagsNoPicker ColorPickerFlags = 1 << 2
	// ColorPickerFlagsNoOptions disables toggling options menu when right-clicking on inputs/small preview.
	ColorPickerFlagsNoOptions ColorPickerFlags = 1 << 3
	// ColorPickerFlagsNoAlpha ignoreс Alpha component (read 3 components from the input pointer).
	ColorPickerFlagsNoAlpha ColorPickerFlags = 1 << 1
	// ColorPickerFlagsNoSmallPreview disables colored square preview next to the inputs. (e.g. to show only the inputs).
	ColorPickerFlagsNoSmallPreview ColorPickerFlags = 1 << 4
	// ColorPickerFlagsNoInputs disables inputs sliders/text widgets (e.g. to show only the small preview colored square).
	ColorPickerFlagsNoInputs ColorPickerFlags = 1 << 5
	// ColorPickerFlagsNoTooltip disables tooltip when hovering the preview.
	ColorPickerFlagsNoTooltip ColorPickerFlags = 1 << 6
	// ColorPickerFlagsNoLabel disables display of inline text label (the label is still forwarded to the tooltip and picker).
	ColorPickerFlagsNoLabel ColorPickerFlags = 1 << 7
	// ColorPickerFlagsNoSidePreview disables bigger color preview on right side of the picker, use small colored square preview instead.
	ColorPickerFlagsNoSidePreview ColorPickerFlags = 1 << 8

	// User Options (right-click on widget to change some of them). You can set application defaults using SetColorEditOptions().
	// The idea is that you probably don't want to override them in most of your calls, let the user choose and/or call
	// SetColorPickerOptions() during startup.

	// ColorPickerFlagsAlphaBar shows vertical alpha bar/gradient in picker.
	ColorPickerFlagsAlphaBar ColorPickerFlags = 1 << 16
	// ColorPickerFlagsAlphaPreview displays preview as a transparent color over a checkerboard, instead of opaque.
	ColorPickerFlagsAlphaPreview ColorPickerFlags = 1 << 17
	// ColorPickerFlagsAlphaPreviewHalf displays half opaque / half checkerboard, instead of opaque.
	ColorPickerFlagsAlphaPreviewHalf ColorPickerFlags = 1 << 18
	// ColorPickerFlagsRGB sets the format as RGB.
	ColorPickerFlagsRGB ColorPickerFlags = 1 << 20
	// ColorPickerFlagsHSV sets the format as HSV.
	ColorPickerFlagsHSV ColorPickerFlags = 1 << 21
	// ColorPickerFlagsHEX sets the format as HEX.
	ColorPickerFlagsHEX ColorPickerFlags = 1 << 22
	// ColorPickerFlagsUint8 _display_ values formatted as 0..255.
	ColorPickerFlagsUint8 ColorPickerFlags = 1 << 23
	// ColorPickerFlagsFloat _display_ values formatted as 0.0f..1.0f floats instead of 0..255 integers. No round-trip of value via integers.
	ColorPickerFlagsFloat ColorPickerFlags = 1 << 24
	// ColorPickerFlagsPickerHueBar bar for Hue, rectangle for Sat/Value.
	ColorPickerFlagsPickerHueBar ColorPickerFlags = 1 << 25
	// ColorPickerFlagsPickerHueWheel wheel for Hue, triangle for Sat/Value.
	ColorPickerFlagsPickerHueWheel ColorPickerFlags = 1 << 26
	// ColorPickerFlagsInputRGB enables input and output data in RGB format.
	ColorPickerFlagsInputRGB ColorPickerFlags = 1 << 27
	// ColorPickerFlagsInputHSV enables input and output data in HSV format.
	ColorPickerFlagsInputHSV ColorPickerFlags = 1 << 28
)

// ColorPicker3 calls ColorPicker3V(label, col, 0).
func ColorPicker3(label string, col *[3]float32) bool {
	return ColorPicker3V(label, col, 0)
}

// ColorPicker3V will show directly a color picker control for editing a color in 3D vector (rgb format).
func ColorPicker3V(label string, col *[3]float32, flags ColorPickerFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	ccol := (*C.float)(&col[0])
	return C.iggColorPicker3(labelArg, ccol, C.int(flags)) != 0
}

// ColorPicker4 calls ColorPicker4V(label, col, 0).
func ColorPicker4(label string, col *[4]float32) bool {
	return ColorPicker4V(label, col, 0)
}

// ColorPicker4V will show directly a color picker control for editing a color in 4D vector (rgba format).
func ColorPicker4V(label string, col *[4]float32, flags ColorPickerFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	ccol := (*C.float)(&col[0])
	return C.iggColorPicker4(labelArg, ccol, C.int(flags)) != 0
}
