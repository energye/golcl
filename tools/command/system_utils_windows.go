//go:build windows
// +build windows

package command

import "syscall"

func HideWindow(bool bool) *syscall.SysProcAttr {
	return &syscall.SysProcAttr{HideWindow: bool}
}
