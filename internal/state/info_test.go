package state_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pkorobeinikov/multipass-compose/internal/state"
)

func TestParseInfo(t *testing.T) {
	t.Run(`positive`, func(t *testing.T) {
		tests := []struct {
			Name         string
			Given        []byte
			ExpectedInfo *state.Info
		}{
			{
				Name:  "multipass-info",
				Given: readFile("testdata/multipass-info.json"),
				ExpectedInfo: &state.Info{
					Info: map[string]state.Service{
						"instance-1": {
							Ipv4: []string{
								"192.168.64.39",
							},
							State: "Running",
						},
					},
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.Name, func(t *testing.T) {
				actualInfo, err := state.ParseInfo(tt.Given)

				assert.NoError(t, err)
				assert.Contains(t, actualInfo.Info, "instance-1")
				assert.Equal(t, tt.ExpectedInfo.Info["instance-1"].Ipv4, actualInfo.Info["instance-1"].Ipv4)
				assert.Equal(t, tt.ExpectedInfo.Info["instance-1"].State, actualInfo.Info["instance-1"].State)
			})
		}
	})
}

func readFile(path string) []byte {
	b, _ := os.ReadFile(path)

	return b
}
