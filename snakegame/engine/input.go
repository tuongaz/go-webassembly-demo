package engine

import "encoding/json"

const (
	ChangeDirection = "change_direction"
)

type Input struct {
	Type string         `json:"type"`
	Data map[string]any `json:"data"`
}

func (e Input) JSON() string {
	body, _ := json.Marshal(e)
	return string(body)
}
