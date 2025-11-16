//go:build windows
// +build windows

package zlog

import (
	"errors"
	"os"
	"syscall"
)

func getStats(info os.FileInfo) (interface{}, error) {
	// 这里是Windows系统的实现
	// 使用syscall.Win32FileAttributeData
	attr, ok := info.Sys().(*syscall.Win32FileAttributeData)
	if !ok {
		return 0, errors.New("failed to get Win32FileAttributeData")
	}
	return attr, nil
}
