//go:build windows

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"syscall"
)

// The duration of the time-out period, in milliseconds. If the message is a broadcast message,
// each window can use the full time-out period:
// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-sendmessagetimeouta
const TIMEOUT_MS = 5000

// https://goreleaser.com/cookbooks/using-main.version/
var (
	name    string
	version string
	date    string
	commit  string
)

// flags
type Config struct {
	help    bool
	version bool
}

func initFlags() *Config {
	cfg := &Config{}
	flag.BoolVar(&cfg.help, "?", false, "")
	flag.BoolVar(&cfg.help, "help", false, "displays this help message")
	flag.BoolVar(&cfg.version, "v", false, "")
	flag.BoolVar(&cfg.version, "version", false, "print version and exit")
	return cfg
}

func main() {
	log.SetFlags(0)
	cfg := initFlags()
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: "+name+` [OPTIONS]

Refresh environment variables from the Windows registry.

OPTIONS:

  -?, --help
          display this help message
  -v, --version
          print version and exit`)
	}
	flag.Parse()

	if flag.Arg(0) == "version" || cfg.version {
		fmt.Printf("%s %s, built on %s (commit: %s)\n", name, version, date, commit)
		return
	}

	if cfg.help {
		flag.Usage()
		return
	}

	lParam, _ := syscall.UTF16PtrFromString("Environment")

	// When an application sends this message, wParam must be NULL:
	// https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-settingchange
	ret := SendMessageTimeout(HWND_BROADCAST, WM_SETTINGCHANGE, nil,
		lParam, SMTO_NORMAL|SMTO_ABORTIFHUNG, TIMEOUT_MS)

	// If the function succeeds, the return value is nonzero
	if ret == 0 {
		fmt.Println("Refresh: Error")
		os.Exit(1)
	} else {
		fmt.Println("Refresh: Success")
	}
}
