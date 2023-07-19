package food

import (
	"math/rand"
	"snakegame/coord"
	"time"
)

func NewDefaultFoodProvider(width, height int) *DefaultFoodProvider {
	return &DefaultFoodProvider{
		width:  width,
		height: height,
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

type DefaultFoodProvider struct {
	width  int
	height int
	rand   *rand.Rand
}

func (d *DefaultFoodProvider) CreateFood() Food {
	foods := []rune{'🍟', '🍔', '🍕'}
	return Food{
		Coord: coord.Coord{
			X: d.rand.Intn(d.width-5) + 2,
			Y: d.rand.Intn(d.height-5) + 2,
		},
		Image: foods[d.rand.Intn(len(foods))],
	}
}
