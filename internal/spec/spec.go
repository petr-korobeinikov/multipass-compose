package spec

import (
	"bytes"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/pkorobeinikov/multipass-compose/internal/cfg"
)

type (
	Spec struct {
		Services map[string]Service `yaml:"services"`
	}

	Service struct {
		Image     string `yaml:"image"`
		CPUs      string `yaml:"cpus,omitempty"`
		Mem       string `yaml:"mem,omitempty"`
		Disk      string `yaml:"disk,omitempty"`
		CloudInit string `yaml:"cloud-init,omitempty"`
	}
)

func Load(specFileName string) (*Spec, error) {
	var s Spec

	f, err := os.ReadFile(specFileName)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(f, &s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func Encode(s *Spec) ([]byte, error) {
	var b bytes.Buffer

	enc := yaml.NewEncoder(&b)
	defer enc.Close()

	enc.SetIndent(cfg.DefaultYamlIndent)

	if err := enc.Encode(s); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
