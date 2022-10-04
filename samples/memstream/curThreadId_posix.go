//go:build !windows
// +build !windows

package main

import (
	"syscall"
)

func GetCurrentThreadId() int {
	return syscall.Gettid()
}
