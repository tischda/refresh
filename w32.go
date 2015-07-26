package main

// Cgo recognizes the comment above the import statement... these are used as
// a header when compiling the C parts of the package. In this case those
// lines are just a single #include statement, but they can be almost any C code.
import (
	// #include <wtypes.h>
	// #include <winable.h>
	"C"
	"log"
	"syscall"
	"time"
	"unsafe"
)

type HWND uintptr

var (
	moduser32 = syscall.NewLazyDLL("user32.dll")

	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms633499(v=vs.85).aspx
	procFindWindowW = moduser32.NewProc("FindWindowW")

	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms646310(v=vs.85).aspx
	procSendInput = moduser32.NewProc("SendInput")
)

const (
	KEYEVENTF_KEYDOWN = 0
	KEYEVENTF_KEYUP   = 0x0002
	SENDKEYS_DELAY    = 100 * time.Millisecond
)

// extract copied from "github.com/AllenDang/w32"

// Virtual-Key Codes
const (
	VK_RETURN = 0x0D
)

// Inspired by http://play.golang.org/p/kwfYDhhiqk
func sendKey(vk uint16) {
	var inputs []INPUT
	inputs = append(inputs, INPUT{
		Type: INPUT_KEYBOARD,
		Ki:   keyPress(vk, KEYEVENTF_KEYDOWN),
	})
	inputs = append(inputs, INPUT{
		Type: INPUT_KEYBOARD,
		Ki:   keyPress(vk, KEYEVENTF_KEYUP),
	})
	SendInput(inputs)
}

func keyPress(vk uint16, event uint32) KEYBDINPUT {
	return KEYBDINPUT{
		WVk:         vk,
		WScan:       0,
		DwFlags:     event,
		Time:        0,
		DwExtraInfo: 0,
	}
}

func SendInput(inputs []INPUT) uint32 {
	ret, _, _ := procSendInput.Call(
		uintptr(len(inputs)),
		uintptr(unsafe.Pointer(&inputs[0])),
		uintptr(unsafe.Sizeof(C.INPUT{})),
	)
	return uint32(ret)
}

const (
	INPUT_KEYBOARD = 1
)

// http://msdn.microsoft.com/en-us/library/windows/desktop/ms646270(v=vs.85).aspx
type INPUT struct {
	Type uint32
	Ki   KEYBDINPUT
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/ms646271(v=vs.85).aspx
type KEYBDINPUT struct {
	WVk         uint16
	WScan       uint16
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
}

// shorter version of: http://play.golang.org/p/kwfYDhhiqk
// see: https://github.com/vevix/twitch-plays/blob/master/win32/win32.go#L23
func findWindow(title string) HWND {
	ret, _, _ := procFindWindowW.Call(0, uintptr(unsafe.Pointer(StringToUTF16Ptr(title))))
	if ret == 0 {
		log.Fatalln("Cannot find window:", title)
	}
	return HWND(ret)
}

// https://golang.org/src/syscall/syscall_windows.go
// syscall.StringToUTF16Ptr is deprecated, this is our own:
func StringToUTF16Ptr(s string) *uint16 {
	a, err := syscall.UTF16FromString(s)
	if err != nil {
		log.Fatalln("syscall: string with NUL passed to StringToUTF16")
	}
	return &a[0]
}
