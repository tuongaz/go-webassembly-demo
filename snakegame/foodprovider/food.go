package foodprovider

import (
	"context"
	_ "embed"
	"encoding/json"
	"log"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"

	"snakegame/coord"
	"snakegame/engine/food"
)

//go:embed target/wasm32-wasi/debug/foodprovider.wasm
var foodWasm []byte

func New(width, height int) *ExternalFoodProvider {
	ctx := context.Background()
	r := wazero.NewRuntime(ctx)
	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	mod, err := r.Instantiate(ctx, foodWasm)
	if err != nil {
		log.Panicf("failed to instantiate module: %v", err)
	}

	return &ExternalFoodProvider{
		width:  uint64(width),
		height: uint64(height),
		module: mod,
	}
}

type ExternalFoodProvider struct {
	width  uint64
	height uint64
	module api.Module
}

func (d *ExternalFoodProvider) CreateFood() food.Food {
	// Choose the context to use for function calls.
	ctx := context.Background()

	results, err := d.module.ExportedFunction("create_food").Call(ctx, d.width, d.height)
	if err != nil {
		log.Panicf("failed to call: %v", err)
	}

	ptr := results[0] >> 32
	length := results[0] & 0xFFFFFFFF

	// Read data from memory
	data, _ := d.module.Memory().Read(uint32(ptr), uint32(length))

	f := map[string]any{}
	if err := json.Unmarshal(data, &f); err != nil {
		panic(err)
	}

	x := f["coordination"].(map[string]any)["x"].(float64)
	y := f["coordination"].(map[string]any)["y"].(float64)

	return food.Food{
		Coord: coord.Coord{
			X: int(x),
			Y: int(y),
		},
		Image: []rune(f["image"].(string))[0],
	}
}
