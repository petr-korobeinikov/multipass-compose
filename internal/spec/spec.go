package spec

import (
	"os"

	"gopkg.in/yaml.v3"
)

type (
	Spec struct {
		Services map[string]Service `yaml:"services"`
	}

	Service struct {
		Image string `yaml:"image"`
		CPUs  string `yaml:"cpus"`
		Mem   string `yaml:"mem"`
		Disk  string `yaml:"disk"`
	}
)

func Load(path string) (*Spec, error) {
	var s Spec

	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(f, &s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
