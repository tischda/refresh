// +build windows

package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"syscall"
	"time"
	"unsafe"

	"github.com/AllenDang/w32"
)

// http://technosophos.com/2014/06/11/compile-time-string-in-go.html
// go build -ldflags "-x main.version $(git describe --tags)"
var version string

var (
	moduser32 = syscall.NewLazyDLL("user32.dll")

	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms633499(v=vs.85).aspx
	procFindWindowW = moduser32.NewProc("FindWindowW")
)

const (
	KEYEVENTF_KEYDOWN = 0
	KEYEVENTF_KEYUP   = 0x0002
	SENDKEYS_DELAY    = 100 * time.Millisecond
)

var showVersion bool
var panelTitle string

func init() {
	flag.BoolVar(&showVersion, "version", false, "print version and exit")
	flag.StringVar(&panelTitle, "title", "Environment Variables", "localized version of title")
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	if showVersion {
		fmt.Println("refresh version", version)
	} else {
		cmd := exec.Command("cmd", "/c", "start rundll32 sysdm.cpl,EditEnvironmentVariables")
		err := cmd.Run()
		if err != nil {
			log.Fatal("exec.Command:", err)
		}
		time.Sleep(SENDKEYS_DELAY)
		findWindow(panelTitle)
		sendkey(w32.VK_RETURN)
	}
}

// Inspired by http://play.golang.org/p/kwfYDhhiqk
func sendkey(vk uint16) {
	var inputs []w32.INPUT
	inputs = append(inputs, w32.INPUT{
		Type: w32.INPUT_KEYBOARD,
		Ki:   keyPress(vk, KEYEVENTF_KEYDOWN),
	})
	inputs = append(inputs, w32.INPUT{
		Type: w32.INPUT_KEYBOARD,
		Ki:   keyPress(vk, KEYEVENTF_KEYUP),
	})
	w32.SendInput(inputs)
}

func keyPress(vk uint16, event uint32) w32.KEYBDINPUT {
	return w32.KEYBDINPUT{
		WVk:         vk,
		WScan:       0,
		DwFlags:     event,
		Time:        0,
		DwExtraInfo: 0,
	}
}

// shorter version of: http://play.golang.org/p/kwfYDhhiqk
// see: https://github.com/vevix/twitch-plays/blob/master/win32/win32.go#L23
func findWindow(title string) w32.HWND {
	ret, _, _ := procFindWindowW.Call(0, uintptr(unsafe.Pointer(StringToUTF16Ptr(title))))
	if ret == 0 {
		log.Fatalln("Cannot find window:", title)
	}
	return w32.HWND(ret)
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
