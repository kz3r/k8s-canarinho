package settings

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type YamlConf struct {
	Conf struct {
		Namespaces             []string
		CanaryMaxTimeInMinutes int32 `yaml:"canaryMaxTimeInMinutes"`
	}
}

func ReadConf(filename string) (*YamlConf, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	conf := &YamlConf{}
	err = yaml.Unmarshal(buf, conf)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %w", filename, err)
	}

	return conf, err
}
