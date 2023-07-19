// Some code related to render the snake and frame in terminal was inspired by
// https://github.com/Simant-Thapa-Magar/snake-go

package term

import (
	"fmt"
	"os"

	"snakegame/coord"
	"snakegame/engine"
	"snakegame/engine/food"

	"github.com/gdamore/tcell/v2"
)

const (
	FRAME_BORDER_THICKNESS    = 1
	FRAME_BORDER_VERTICAL     = '║'
	FRAME_BORDER_HORIZONTAL   = '═'
	FRAME_BORDER_TOP_LEFT     = '╔'
	FRAME_BORDER_TOP_RIGHT    = '╗'
	FRAME_BORDER_BOTTOM_RIGHT = '╝'
	FRAME_BORDER_BOTTOM_LEFT  = '╚'
)

var _ engine.Platform = &Term{}
var eventCh = make(chan engine.Input)
var coordinatesToClear []coord.Coord

type Term struct {
	screen tcell.Screen
}

func New() (*Term, error) {
	screen, err := initScreen()
	if err != nil {
		return nil, err
	}

	t := &Term{
		screen: screen,
	}
	t.waitForUserInputs()

	return t, nil
}

func (t *Term) GameReady(board *engine.Board) {
	t.screen.Clear()
	t.displayFrame(board.Width, board.Height)
	t.render(board)
}

func (t *Term) Input() chan engine.Input {
	return eventCh
}

func (t *Term) Output(board *engine.Board) {
	t.render(board)
}

func initScreen() (tcell.Screen, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	if err = screen.Init(); err != nil {
		return nil, err
	}

	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	screen.SetStyle(defStyle)

	return screen, nil
}

func (t *Term) waitForUserInputs() {
	userInput := t.readUserInput()
	go func() {
		for {
			key := t.getUserInput(userInput)
			if key == "Rune[q]" || key == "Ctrl+C" {
				t.screen.Fini()
				os.Exit(0)
			} else {
				if key == "Up" {
					t.changeDirection(string(engine.UP))
				} else if key == "Down" {
					t.changeDirection(string(engine.DOWN))
				} else if key == "Left" {
					t.changeDirection(string(engine.LEFT))
				} else if key == "Right" {
					t.changeDirection(string(engine.RIGHT))
				}
			}
		}
	}()
}

func (t *Term) drawSnake(b []coord.Coord) {
	var body = make([]coord.Coord, len(b))
	copy(body, b)
	for i, cord := range b {
		body[i] = t.transformCord(cord)
	}

	style := tcell.StyleDefault.Foreground(tcell.ColorDarkGreen.TrueColor())
	for _, cord := range body {
		t.print(cord.X, cord.Y, 1, 1, style, 0x2588)
	}
	coordinatesToClear = append(coordinatesToClear, body[0])
	t.screen.Show()
}

func (t *Term) transformCord(cord coord.Coord) coord.Coord {
	return coord.Coord{
		X: cord.X + FRAME_BORDER_THICKNESS,
		Y: cord.Y + FRAME_BORDER_THICKNESS,
	}
}

func (t *Term) drawFood(foods []food.Food) {
	style := tcell.StyleDefault.Foreground(tcell.ColorDarkRed.TrueColor())
	for _, f := range foods {
		pos := t.transformCord(coord.Coord{
			X: f.Coord.X,
			Y: f.Coord.Y,
		})

		t.print(pos.X, pos.Y, 1, 1, style, f.Image)
		coordinatesToClear = append(coordinatesToClear, pos)
	}
	t.screen.Show()
}

func (t *Term) drawPoints(points int) {
	content := fmt.Sprintf("Points: %v", points)

	for i := 0; i < len(content); i++ {
		t.print(i+1, 1, 1, 1, tcell.StyleDefault, rune(content[i]))
		coordinatesToClear = append(coordinatesToClear, coord.Coord{X: i + 1, Y: 1})
	}
	t.screen.Show()
}

func (t *Term) clear() {
	for _, coordinate := range coordinatesToClear {
		t.print(coordinate.X, coordinate.Y, 1, 1, tcell.StyleDefault, ' ')
	}
}

func (t *Term) changeDirection(direction string) {
	eventCh <- engine.Input{
		Type: engine.ChangeDirection,
		Data: map[string]any{
			"direction": direction,
		},
	}
}

func (t *Term) print(x, y, w, h int, style tcell.Style, char rune) {
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			t.screen.SetContent(x+i, y+j, char, nil, style)
		}
	}
}

func (t *Term) render(board *engine.Board) {
	t.clear()

	t.drawSnake(board.Snake.Body)
	t.drawFood(board.Foods)
	t.drawPoints(board.Snake.Points)
}

func (t *Term) displayFrame(width, height int) {
	t.printUnfilledRectangle(
		0,
		0,
		width+3*FRAME_BORDER_THICKNESS,
		height+3*FRAME_BORDER_THICKNESS,
		FRAME_BORDER_THICKNESS,
		FRAME_BORDER_HORIZONTAL,
		FRAME_BORDER_VERTICAL,
		FRAME_BORDER_TOP_LEFT,
		FRAME_BORDER_TOP_RIGHT,
		FRAME_BORDER_BOTTOM_RIGHT,
		FRAME_BORDER_BOTTOM_LEFT,
	)
	t.screen.Show()
}

func (t *Term) printUnfilledRectangle(xOrigin, yOrigin, width, height, borderThickness int, horizontalOutline, verticalOutline, topLeftOutline, topRightOutline, bottomRightOutline, bottomLeftOutline rune) {
	var upperBorder, lowerBorder rune
	verticalBorder := verticalOutline
	for i := 0; i < width; i++ {
		if i == 0 {
			upperBorder = topLeftOutline
			lowerBorder = bottomLeftOutline
		} else if i == width-1 {
			upperBorder = topRightOutline
			lowerBorder = bottomRightOutline
		} else {
			upperBorder = horizontalOutline
			lowerBorder = horizontalOutline
		}
		t.print(xOrigin+i, yOrigin, borderThickness, borderThickness, tcell.StyleDefault, upperBorder)
		t.print(xOrigin+i, yOrigin+height-1, borderThickness, borderThickness, tcell.StyleDefault, lowerBorder)
	}

	for i := 1; i < height-1; i++ {
		t.print(xOrigin, yOrigin+i, borderThickness, borderThickness, tcell.StyleDefault, verticalBorder)
		t.print(xOrigin+width-1, yOrigin+i, borderThickness, borderThickness, tcell.StyleDefault, verticalBorder)
	}
}

func (t *Term) readUserInput() chan string {
	userInput := make(chan string)
	go func() {
		for {
			switch ev := t.screen.PollEvent().(type) {
			case *tcell.EventKey:
				userInput <- ev.Name()
			}
		}
	}()
	return userInput
}

func (t *Term) getUserInput(userInput chan string) string {
	select {
	case key := <-userInput:
		return key
	}
}
