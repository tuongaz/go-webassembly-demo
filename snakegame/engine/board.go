package engine

import (
	"encoding/json"
	"snakegame/coord"

	"snakegame/engine/food"
)

type Board struct {
	Width        int         `json:"width"`
	Height       int         `json:"height"`
	Snake        *Snake      `json:"snake"`
	Foods        []food.Food `json:"foods"`
	foodProvider food.Provider
}

func (b *Board) JSON() string {
	body, _ := json.Marshal(b)
	return string(body)
}

func newBoard(width, height int, foodProvider food.Provider) *Board {
	b := &Board{
		Width:  width,
		Height: height,
		Snake: newSnake(coord.Coord{
			X: width / 2,
			Y: height / 2,
		}),
		foodProvider: foodProvider,
	}
	b.generateFood()
	return b
}

func (b *Board) moveSnake() error {
	if b.snakeHitBorder() {
		return b.Snake.die("snake left Board")
	}

	if err := b.Snake.move(); err != nil {
		return err
	}

	b.handleSnakeFoods()

	return nil
}

func (b *Board) handleSnakeFoods() {
	head := b.Snake.head()
	for i, f := range b.Foods {
		if f.Coord.X == head.X && f.Coord.Y == head.Y {
			b.Snake.addPoints(1)
			b.Snake.grow()
			b.Foods = append(b.Foods[:i], b.Foods[i+1:]...)
			b.generateFood()
		}
	}
}

func (b *Board) snakeHitBorder() bool {
	h := b.Snake.head()
	return h.X >= b.Width || h.Y >= b.Height || h.X <= 0 || h.Y <= 0
}

func (b *Board) generateFood() {
	b.Foods = append(b.Foods, b.foodProvider.CreateFood())
}
