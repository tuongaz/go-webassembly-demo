package main

import (
	"snakegame/engine"
	"snakegame/engine/food"
	"snakegame/platform/web"
)

const (
	gameSize = 15
)

func main() {
	webPlatform := web.New()

	g := engine.New(
		gameSize,
		gameSize,
		webPlatform,
		food.NewDefaultFoodProvider(gameSize, gameSize),
		//foodprovider.New(gameSize, gameSize),
	)
	g.Start()
}
