package spec

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/petr-korobeinikov/multipass-compose/internal/cfg"
	"github.com/petr-korobeinikov/multipass-compose/internal/fsext"
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

var (
	ErrSpecFileNotFound = errors.New("spec file not found")
)

func Load(specFileName string) (*Spec, error) {
	var (
		s            Spec
		specFilePath string
	)

	if fsext.FileExists(specFileName) {
		specFilePath = specFileName
	} else {
		var parentDirsTillRoot []string

		dir := ".."
		for {
			parentDirsTillRoot = append(parentDirsTillRoot, dir)

			dir = filepath.Join("..", dir)

			abs, err := filepath.Abs(dir)
			if err != nil {
				return nil, err
			}
			if abs == string(filepath.Separator) {
				break
			}
		}

		found := false
		for _, d := range parentDirsTillRoot {
			specFilePath = filepath.Join(d, specFileName)
			if fsext.FileExists(specFilePath) {
				found = true

				break
			}
		}
		if !found {
			return nil, ErrSpecFileNotFound
		}
	}

	f, err := os.ReadFile(specFilePath)
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
