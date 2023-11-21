package spec_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/petr-korobeinikov/multipass-compose/internal/spec"
)

func TestLoad(t *testing.T) {
	t.Run(`positive`, func(t *testing.T) {
		tests := []struct {
			Name         string
			GivenPath    string
			ExpectedSpec *spec.Spec
		}{
			{
				Name:      "minimal",
				GivenPath: "testdata/multipass-compose-minimal.yaml",
				ExpectedSpec: &spec.Spec{
					Services: map[string]spec.Service{
						"instance-1": {
							Image: "focal",
						},
					}},
			},
			{
				Name:      "minimal-vm-prop",
				GivenPath: "testdata/multipass-compose-minimal-vm-prop.yaml",
				ExpectedSpec: &spec.Spec{
					Services: map[string]spec.Service{
						"instance-1": {
							Image: "focal",
							CPUs:  "4",
							Mem:   "6G",
							Disk:  "42G",
						},
					},
				},
			},
			{
				Name:      "realistic",
				GivenPath: "testdata/multipass-compose-realistic.yaml",
				ExpectedSpec: &spec.Spec{
					Services: map[string]spec.Service{
						"lb": {
							Image: "focal",
						},
						"instance-1": {
							Image: "focal",
						},
						"instance-2": {
							Image: "focal",
						},
						"instance-3": {
							Image: "focal",
						},
					}},
			},
		}

		for _, tt := range tests {
			t.Run(tt.Name, func(t *testing.T) {
				actualSpec, err := spec.Load(tt.GivenPath)

				assert.NoError(t, err)
				assert.Equal(t, tt.ExpectedSpec, actualSpec)
			})
		}
	})
}
