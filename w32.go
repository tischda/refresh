package main

// Cgo recognizes the comment above the import statement... these are used as
// a header when compiling the C parts of the package. In this case those
// lines are just a single #include statement, but they can be almost any C code.
import (
	// #include <wtypes.h>
	// #include <winable.h>
	"C"
	"unsafe"
)

type HWND uintptr

// extract copied from "github.com/AllenDang/w32"

// Virtual-Key Codes
const (
	VK_RETURN = 0x0D
)

func SendInput(inputs []INPUT) uint32 {
	var validInputs []C.INPUT

	for _, oneInput := range inputs {
		input := C.INPUT{_type: C.DWORD(oneInput.Type)}

		switch oneInput.Type {
		case INPUT_MOUSE:
			(*MouseInput)(unsafe.Pointer(&input)).mi = oneInput.Mi
		case INPUT_KEYBOARD:
			(*KbdInput)(unsafe.Pointer(&input)).ki = oneInput.Ki
		case INPUT_HARDWARE:
			(*HardwareInput)(unsafe.Pointer(&input)).hi = oneInput.Hi
		default:
			panic("unkown type")
		}

		validInputs = append(validInputs, input)
	}

	ret, _, _ := procSendInput.Call(
		uintptr(len(validInputs)),
		uintptr(unsafe.Pointer(&validInputs[0])),
		uintptr(unsafe.Sizeof(C.INPUT{})),
	)
	return uint32(ret)
}

const (
	INPUT_MOUSE    = 0
	INPUT_KEYBOARD = 1
	INPUT_HARDWARE = 2
)

// http://msdn.microsoft.com/en-us/library/windows/desktop/ms646270(v=vs.85).aspx
type INPUT struct {
	Type uint32
	Mi   MOUSEINPUT
	Ki   KEYBDINPUT
	Hi   HARDWAREINPUT
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/ms646273(v=vs.85).aspx
type MOUSEINPUT struct {
	Dx          int32
	Dy          int32
	MouseData   uint32
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/ms646271(v=vs.85).aspx
type KEYBDINPUT struct {
	WVk         uint16
	WScan       uint16
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/ms646269(v=vs.85).aspx
type HARDWAREINPUT struct {
	UMsg    uint32
	WParamL uint16
	WParamH uint16
}

type KbdInput struct {
	typ uint32
	ki  KEYBDINPUT
}

type MouseInput struct {
	typ uint32
	mi  MOUSEINPUT
}

type HardwareInput struct {
	typ uint32
	hi  HARDWAREINPUT
}
