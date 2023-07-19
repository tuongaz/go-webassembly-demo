package engine

import (
	"context"
	"time"

	"snakegame/engine/food"
)

var (
	snakeInitialSpeed = 150
)

type Game struct {
	width        int
	height       int
	platform     Platform
	foodProvider food.Provider
}

func New(width, height int, platform Platform, foodProvider food.Provider) *Game {
	return &Game{
		width:        width,
		height:       height,
		platform:     platform,
		foodProvider: foodProvider,
	}
}

func (g *Game) Start() {
	errCh := make(chan error)
	ctx, stopGameFn := context.WithCancel(context.Background())

	board := newBoard(g.width, g.height, g.foodProvider)

	// Notify game platform to prepare for new game
	g.platform.GameReady(board)
	g.waitForPlayerToStart()

	go g.moveSnake(board, stopGameFn, errCh)
	go g.handleInput(ctx, board)

	<-errCh
	g.Start()
}

func (g *Game) waitForPlayerToStart() {
	// Wait for start signal from the platform
	<-g.platform.Input()
}

func (g *Game) moveSnake(board *Board, stopGameFn context.CancelFunc, errCh chan error) {
	for {
		if err := board.moveSnake(); err != nil {
			stopGameFn()
			errCh <- err
			return
		}

		// notify platform
		g.platform.Output(board)

		time.Sleep(g.speed(board.Snake.Points))
	}
}

func (g *Game) handleInput(ctx context.Context, board *Board) {
	for {
		select {
		// receive inputs from platform
		case ev := <-g.platform.Input():
			if ev.Type == ChangeDirection {
				board.Snake.changeDirection(Direction(ev.Data["direction"].(string)))
			}
		case <-ctx.Done():
			return
		}
	}
}

func (g *Game) speed(points int) time.Duration {
	ms := snakeInitialSpeed - points
	return time.Duration(ms) * time.Millisecond
}
