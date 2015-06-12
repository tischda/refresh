// +build windows

package main

// Inspired from http://play.golang.org/p/kwfYDhhiqk

import (
	"flag"
	"fmt"
	"github.com/AllenDang/w32"
	"log"
	"os/exec"
	"syscall"
	"time"
	"unsafe"
)

const version string = "1.0.0"

var (
	moduser32 = syscall.NewLazyDLL("user32.dll")

	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms633499(v=vs.85).aspx
	procFindWindowW = moduser32.NewProc("FindWindowW")

	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms633539(v=vs.85).aspx
	procSetForegroundWindow = moduser32.NewProc("SetForegroundWindow")
)

const (
	KEYEVENTF_KEYDOWN = 0
	KEYEVENTF_KEYUP   = 0x0002
)

type HWND uintptr

func main() {
	showVersion := flag.Bool("version", false, "print version and exit")

	// configure logging
	log.SetFlags(0)
	flag.Parse()

	if *showVersion {
		fmt.Println("refresh version", version)
		return
	}

	cmd := exec.Command("cmd", "/c", "start rundll32 sysdm.cpl,EditEnvironmentVariables")
	err := cmd.Run()
	if err != nil {
		log.Fatal("exec.Command", err)
	}

	time.Sleep(time.Millisecond * 100)

	hwnd, err := FindWindow("Environment Variables")
	if err != nil {
		log.Fatalln("FindWindow", err)
	}
	ret, _, _ := procSetForegroundWindow.Call(uintptr(hwnd))
	if ret != 1 {
		log.Fatalln("Could not set window to foreground")
	}
	sendkey(w32.VK_RETURN)
}

func sendkey(vk uint16) {
	var inputs []w32.INPUT
	inputs = append(inputs, w32.INPUT{
		Type: w32.INPUT_KEYBOARD,
		Ki: w32.KEYBDINPUT{
			WVk:         vk,
			WScan:       0,
			DwFlags:     KEYEVENTF_KEYDOWN,
			Time:        0,
			DwExtraInfo: 0,
		},
	})

	inputs = append(inputs, w32.INPUT{
		Type: w32.INPUT_KEYBOARD,
		Ki: w32.KEYBDINPUT{
			WVk:         vk,
			WScan:       0,
			DwFlags:     KEYEVENTF_KEYUP,
			Time:        0,
			DwExtraInfo: 0,
		},
	})
	w32.SendInput(inputs)
}

func FindWindow(win string) (ret HWND, err error) {
	lpszWindow := syscall.StringToUTF16Ptr(win)

	r0, _, e1 := syscall.Syscall(procFindWindowW.Addr(), 2, 0, uintptr(unsafe.Pointer(lpszWindow)), 0)
	ret = HWND(r0)
	if ret == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}
