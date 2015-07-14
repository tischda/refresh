// +build windows

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"
)

// Inspired by https://talks.golang.org/2014/testing.slide#23
func TestVersion(t *testing.T) {

	if os.Getenv("BE_CRASHER") == "1" {
		main()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestVersion", "-version")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")

	// capture output of process execution
	r, w, _ := os.Pipe()
	cmd.Stdout = w
	err := cmd.Run()
	w.Close()

	// check return code
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		t.Fatalf("Exptected exit status 0, but was: %v, ", err)
	}

	// now check that version is displayed
	captured, _ := ioutil.ReadAll(r)
	actual := string(captured)
	expected := fmt.Sprintf("refresh version %s\n", version)

	if !strings.Contains(actual, expected) {
		t.Errorf("Expected: %s, but was: %s", expected, actual)
	}
}
