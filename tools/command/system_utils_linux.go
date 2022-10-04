//go:build linux
// +build linux

package command

import "syscall"

func HideWindow(bool bool) *syscall.SysProcAttr {
	return nil
}
