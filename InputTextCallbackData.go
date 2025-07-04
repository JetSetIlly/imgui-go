package imgui

// #include "wrapper/InputTextCallbackData.h"
import "C"
import (
	"sync"
	"unsafe"
)

// InputTextCallback is called for sharing state of an input field.
// By default, the callback should return 0.
type InputTextCallback func(InputTextCallbackData) int32

type inputTextState struct {
	buf *stringBuffer

	key      C.int
	callback InputTextCallback
}

var inputTextStates = make(map[C.int]*inputTextState)
var inputTextStatesMutex sync.Mutex

func newInputTextState(text string, cb InputTextCallback) *inputTextState {
	state := &inputTextState{}
	state.buf = newStringBuffer(text)
	state.callback = cb
	state.register()
	return state
}

func (state *inputTextState) register() {
	inputTextStatesMutex.Lock()
	defer inputTextStatesMutex.Unlock()
	key := C.int(len(inputTextStates) + 1)
	for _, existing := inputTextStates[key]; existing; _, existing = inputTextStates[key] {
		key++
	}
	state.key = key
	inputTextStates[key] = state
}

func (state *inputTextState) release() {
	state.buf.free()

	if state.key != 0 {
		inputTextStatesMutex.Lock()
		defer inputTextStatesMutex.Unlock()
		delete(inputTextStates, state.key)
	}
}

func (state *inputTextState) onCallback(handle C.IggInputTextCallbackData) C.int {
	data := InputTextCallbackData{state: state, handle: handle}
	if data.EventFlag() == InputTextFlagsCallbackResize {
		state.buf.resizeTo(data.bufSize())
		data.setBuf(state.buf.ptr, state.buf.size, data.bufTextLen())
		return 0
	}
	if state.callback == nil {
		return 0
	}
	return C.int(state.callback(data))
}

//export iggInputTextCallback
func iggInputTextCallback(handle C.IggInputTextCallbackData, key C.int) C.int {
	state := iggInputTextStateFor(key)
	return state.onCallback(handle)
}

func iggInputTextStateFor(key C.int) *inputTextState {
	inputTextStatesMutex.Lock()
	defer inputTextStatesMutex.Unlock()
	return inputTextStates[key]
}

// InputTextCallbackData represents the shared state of InputText(), passed as an argument to your callback.
type InputTextCallbackData struct {
	state  *inputTextState
	handle C.IggInputTextCallbackData
}

// EventFlag returns one of the InputTextFlagsCallback* constants to indicate the nature of the callback.
func (data InputTextCallbackData) EventFlag() InputTextFlags {
	return InputTextFlags(C.iggInputTextCallbackDataGetEventFlag(data.handle))
}

// Flags returns the set of flags that the user originally passed to InputText.
func (data InputTextCallbackData) Flags() InputTextFlags {
	return InputTextFlags(C.iggInputTextCallbackDataGetFlags(data.handle)) & ^InputTextFlagsCallbackResize
}

// EventChar returns the current character input. Only valid during CharFilter callback.
func (data InputTextCallbackData) EventChar() rune {
	return rune(C.iggInputTextCallbackDataGetEventChar(data.handle))
}

// SetEventChar overrides what the user entered. Set to zero do drop the current input.
// Returning 1 from the callback also drops the current input.
// Only valid during CharFilter callback.
//
// Note: The internal representation of characters is based on uint16, so less than rune would provide.
func (data InputTextCallbackData) SetEventChar(value rune) {
	C.iggInputTextCallbackDataSetEventChar(data.handle, C.ushort(value))
}

// EventKey returns the currently pressed key. Valid for completion and history callbacks.
func (data InputTextCallbackData) EventKey() ImguiKey {
	return ImguiKey(C.iggInputTextCallbackDataGetEventKey(data.handle))
}

// Buffer returns a view into the current UTF-8 buffer.
// Only during the callbacks of [Completion,History,Always] the current buffer is returned.
// The returned slice is a temporary view into the underlying raw buffer. Do not keep it!
// The underlying memory allocation may even change through a call to InsertBytes().
//
// You may change the buffer through the following ways:
// If the new text has a different (encoded) length, use the functions InsertBytes() and/or DeleteBytes().
// Otherwise you may keep the buffer as is and modify the bytes. If you change the buffer this way directly, mark the buffer
// as modified with MarkBufferModified().
func (data InputTextCallbackData) Buffer() []byte {
	ptr := C.iggInputTextCallbackDataGetBuf(data.handle)
	if ptr == nil {
		return nil
	}
	textLen := data.bufTextLen()
	return ptrToByteSlice(unsafe.Pointer(ptr))[:textLen]
}

// MarkBufferModified indicates that the content of the buffer was modified during a callback.
// Only considered during [Completion,History,Always] callbacks.
func (data InputTextCallbackData) MarkBufferModified() {
	C.iggInputTextCallbackDataMarkBufferModified(data.handle)
}

func (data InputTextCallbackData) setBuf(buf unsafe.Pointer, size, textLen int) {
	C.iggInputTextCallbackDataSetBuf(data.handle, (*C.char)(buf), C.int(size), C.int(textLen))
}

func (data InputTextCallbackData) bufSize() int {
	return int(C.iggInputTextCallbackDataGetBufSize(data.handle))
}

func (data InputTextCallbackData) bufTextLen() int {
	return int(C.iggInputTextCallbackDataGetBufTextLen(data.handle))
}

// DeleteBytes removes the given count of bytes starting at the specified byte offset within the buffer.
// This function can be called during the [Completion,History,Always] callbacks.
// Clears the current selection.
//
// This function ignores the deletion beyond the current buffer length.
// Calling with negative offset or count arguments will panic.
func (data InputTextCallbackData) DeleteBytes(offset, count int) {
	if offset < 0 {
		panic("invalid offset")
	}
	if count < 0 {
		panic("invalid count")
	}
	textLen := data.bufTextLen()
	if offset >= textLen {
		return
	}
	toRemove := count
	available := textLen - offset
	if toRemove > available {
		toRemove = available
	}
	C.iggInputTextCallbackDataDeleteBytes(data.handle, C.int(offset), C.int(toRemove))
}

// InsertBytes inserts the given bytes at given byte offset into the buffer.
// Calling this function may change the underlying buffer allocation.
//
// This function can be called during the [Completion,History,Always] callbacks.
// Clears the current selection.
//
// Calling with an offset outside of the range of the buffer will panic.
func (data InputTextCallbackData) InsertBytes(offset int, bytes []byte) {
	if (offset < 0) || (offset > data.bufTextLen()) {
		panic("invalid offset")
	}
	var bytesPtr *C.char
	byteCount := len(bytes)
	if byteCount > 0 {
		bytesPtr = (*C.char)(unsafe.Pointer(&bytes[0]))
		C.iggInputTextCallbackDataInsertBytes(data.handle, C.int(offset), bytesPtr, C.int(byteCount))
	}
}

// CursorPos returns the byte-offset of the cursor within the buffer.
// Only valid during [Completion,History,Always] callbacks.
func (data InputTextCallbackData) CursorPos() int {
	return int(C.iggInputTextCallbackDataGetCursorPos(data.handle))
}

// SetCursorPos changes the current byte-offset of the cursor within the buffer.
// Only valid during [Completion,History,Always] callbacks.
func (data InputTextCallbackData) SetCursorPos(value int) {
	C.iggInputTextCallbackDataSetCursorPos(data.handle, C.int(value))
}

// SelectionStart returns the byte-offset of the selection start within the buffer.
// Only valid during [Completion,History,Always] callbacks.
func (data InputTextCallbackData) SelectionStart() int {
	return int(C.iggInputTextCallbackDataGetSelectionStart(data.handle))
}

// SetSelectionStart changes the current byte-offset of the selection start within the buffer.
// Only valid during [Completion,History,Always] callbacks.
func (data InputTextCallbackData) SetSelectionStart(value int) {
	C.iggInputTextCallbackDataSetSelectionStart(data.handle, C.int(value))
}

// SelectionEnd returns the byte-offset of the selection end within the buffer.
// Only valid during [Completion,History,Always] callbacks.
func (data InputTextCallbackData) SelectionEnd() int {
	return int(C.iggInputTextCallbackDataGetSelectionEnd(data.handle))
}

// SetSelectionEnd changes the current byte-offset of the selection end within the buffer.
// Only valid during [Completion,History,Always] callbacks.
func (data InputTextCallbackData) SetSelectionEnd(value int) {
	C.iggInputTextCallbackDataSetSelectionEnd(data.handle, C.int(value))
}
