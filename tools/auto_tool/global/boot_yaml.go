package global

import (
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

type ServicesConf struct {
	Tag          string   `yaml:"tag"`
	TagTitleCase string   // 把tag的首字母大写
	TemplateIds  []string `yaml:"templateIds"` // 服务想要启用的功能
}

type BootYaml struct {
	OutPath  string          `yaml:"outpath"`
	Name     string          `yaml:"name"`
	Services []*ServicesConf `yaml:"services"`
}

var BootConfigYaml = &BootYaml{}

func ParseBootYaml(configPath string) error {
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, BootConfigYaml)
	if err != nil {
		return err
	}

	for _, service := range BootConfigYaml.Services {
		service.TagTitleCase = strings.Title(service.Tag)
	}

	return nil
}
