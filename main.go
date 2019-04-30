// +build windows

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var version string
var name = "refresh"

var showVersion bool

func init() {
	flag.BoolVar(&showVersion, "version", false, "print version and exit")
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	if showVersion {
		fmt.Printf("%s version %s\n", name, version)
	} else {
		ret := SendMessageTimeout(HWND_BROADCAST, WM_SETTINGCHANGE, StringToUTF16Ptr(""),
			StringToUTF16Ptr("Environment"), SMTO_NORMAL|SMTO_ABORTIFHUNG, 5000)

		// If the function succeeds, the return value is nonzero
		if ret == 0 {
			fmt.Println("Refresh: Error")
			os.Exit(1)
		} else {
			fmt.Println("Refresh: Success")
		}
	}
}
