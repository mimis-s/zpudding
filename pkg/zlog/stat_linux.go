//go:build linux || darwin
// +build linux darwin

package zlog

import (
	"os"
	"syscall"
)

func getStats(info os.FileInfo) (interface{}, error) {
	// 这里是Unix/Linux系统的实现
	// 使用syscall.Stat_t
	return info.Sys().(*syscall.Stat_t), nil
}
