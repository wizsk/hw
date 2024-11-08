//go:build !windows
// +build !windows

package main

import (
	"os/exec"
	"syscall"
)

func addAtrribute(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true, // This will only be used on Unix-based systems
	}
}
