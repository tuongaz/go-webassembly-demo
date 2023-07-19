package web

import (
	"snakegame/engine"
	"syscall/js"
)

var _ engine.Platform = &Web{}
var eventCh = make(chan engine.Input)

func New() *Web {
	w := &Web{}
	// register this method, it is called from JS
	js.Global().Set("changeDirection", js.FuncOf(w.changeDirection))

	return w
}

type Web struct{}

func (w *Web) GameReady(board *engine.Board) {
	js.Global().Call("gameReady", board.JSON())
}

func (w *Web) Input() chan engine.Input {
	return eventCh
}

func (w *Web) Output(board *engine.Board) {
	js.Global().Call("output", board.JSON())
}

func (w *Web) changeDirection(this js.Value, inputs []js.Value) any {
	//js.Global().Get("console").Call("log", "received change direction from web:", inputs[0].String())
	eventCh <- engine.Input{
		Type: "change_direction",
		Data: map[string]any{
			"direction": inputs[0].String(),
		},
	}
	return nil
}
