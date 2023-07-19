package main

import (
	"snakegame/engine"
	"snakegame/engine/food"
	"snakegame/platform/term"
)

const (
	width  = 30
	height = 15
)

func main() {
	platform, err := term.New()
	if err != nil {
		panic(err)
	}

	g := engine.New(
		width,
		height,
		platform,
		food.NewDefaultFoodProvider(width, height),
		//foodprovider.New(width, height),
	)
	g.Start()
}
