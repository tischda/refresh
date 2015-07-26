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
var name string = "refresh"

var showVersion bool
var panelTitle string
var sendDelay time.Duration

func init() {
	flag.BoolVar(&showVersion, "version", false, "print version and exit")
	flag.StringVar(&panelTitle, "title", "Environment Variables", "localized version of title")
	flag.DurationVar(&sendDelay, "delay", 40*time.Millisecond, "Delays in milliseconds before sending VK_RETURN")
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	if showVersion {
		fmt.Printf("%s version %s\n", name, version)
	} else {
		cmd := exec.Command("cmd", "/c", "start rundll32 sysdm.cpl,EditEnvironmentVariables")
		err := cmd.Run()
		if err != nil {
			log.Fatal("Failed executing command:", err)
		}
		time.Sleep(sendDelay)
		findWindow(panelTitle)
		sendKey(VK_RETURN)
	}
}
