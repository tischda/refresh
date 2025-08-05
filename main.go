//go:build windows
// +build windows

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const PROG_NAME string = "refresh"

// The duration of the time-out period, in milliseconds. If the message is a broadcast message,
// each window can use the full time-out period:
// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-sendmessagetimeouta
const TIMEOUT_MS = 5000

var version string
var flag_version = flag.Bool("version", false, "print version and exit")

func main() {
	log.SetFlags(0)
	flag.Parse()

	if flag.Arg(0) == "version" || *flag_version {
		fmt.Printf("%s version %s\n", PROG_NAME, version)
		return
	}

	// When an application sends this message, wParam must be NULL:
	// https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-settingchange
	ret := SendMessageTimeout(HWND_BROADCAST, WM_SETTINGCHANGE, nil,
		StringToUTF16Ptr("Environment"), SMTO_NORMAL|SMTO_ABORTIFHUNG, TIMEOUT_MS)

	// If the function succeeds, the return value is nonzero
	if ret == 0 {
		fmt.Println("Refresh: Error")
		os.Exit(1)
	} else {
		fmt.Println("Refresh: Success")
	}
}
