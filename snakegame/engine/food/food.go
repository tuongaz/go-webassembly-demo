package food

import (
	"snakegame/coord"
)

type Food struct {
	Coord coord.Coord `json:"coord"`
	Image rune        `json:"image"`
}

type Provider interface {
	CreateFood() Food
}
