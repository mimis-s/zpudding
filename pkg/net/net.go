package net

import (
	"github.com/mimis-s/zpudding/pkg/net/clientConn"
	"github.com/mimis-s/zpudding/pkg/net/http"
	"github.com/mimis-s/zpudding/pkg/net/service"
	"github.com/mimis-s/zpudding/pkg/net/tcp"
)

var mapProtol = make(map[string]service.Service)

func init() {
	mapProtol["tcp"] = new(tcp.Tcp)
	mapProtol["http"] = new(http.Http)

	// mapProtol["udp"] = new(udp.Udp)
}

func InitServer(addr string, sProtocol string, plulgFunc func(clientConn.ClientConn) clientConn.ClientSession) service.Service {

	s := mapProtol[sProtocol]
	s.SetAddr(addr, sProtocol, plulgFunc)
	return s
}
