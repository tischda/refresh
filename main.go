// +build windows

package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"time"
)

// http://technosophos.com/2014/06/11/compile-time-string-in-go.html
// go build -ldflags "-x main.version $(git describe --tags)"
var version string

var showVersion bool
var panelTitle string

const SENDKEYS_DELAY = 100 * time.Millisecond

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
		sendKey(VK_RETURN)
	}
}
