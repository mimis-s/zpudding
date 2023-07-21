package service

import "github.com/mimis-s/zpudding/pkg/net/clientConn"

type Service interface {
	Run() error
	Stop()
	SetAddr(addr, protocol string, newSessionFunc func(clientConn.ClientConn) clientConn.ClientSession)
}
