package main

import (
	"{{.Name}}/src/common/registry"
	"github.com/mimis-s/zpudding/pkg/app"
	{{range $_, $m := .Services}}"{{$.Name}}/src/services/{{$m.Tag}}"
    {{end}}
)

func Boots() []struct {
	bootFunc func() (*app.AppOutSideInfo, error)
} {
	list := []struct {
		bootFunc func() (*app.AppOutSideInfo, error)
	}{
        {{range $_, $m := .Services}}{ {{$m.Tag}}.CreateAppInfo },
        {{end}}
	}

	return list
}

func main() {
	s := registry.NewDefRegistry()
	bootList := Boots()
	for _, appBoot := range bootList {
		info, err := appBoot.bootFunc()
		if err != nil {
			panic(err)
		}
		s.AddAppOutSide(info)
	}

	err := s.Run()
	if err != nil {
		panic(err)
	}
}