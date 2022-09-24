package state

import (
	"encoding/json"
)

type (
	Info struct {
		Errors []interface{}      `json:"errors"`
		Info   map[string]Service `json:"info"`
	}

	Service struct {
		Disks struct {
			Sda1 struct {
				Total string `json:"total"`
				Used  string `json:"used"`
			} `json:"sda1"`
		} `json:"disks"`
		ImageHash    string    `json:"image_hash"`
		ImageRelease string    `json:"image_release"`
		Ipv4         []string  `json:"ipv4"`
		Load         []float64 `json:"load"`
		Memory       struct {
			Total int64 `json:"total"`
			Used  int   `json:"used"`
		} `json:"memory"`
		Mounts struct {
		} `json:"mounts"`
		Release string `json:"release"`
		State   string `json:"state"`
	}
)

func ParseInfo(b []byte) (*Info, error) {
	var i Info

	if err := json.Unmarshal(b, &i); err != nil {
		return nil, err
	}

	return &i, nil
}
