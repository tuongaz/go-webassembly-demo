package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

func main() {
	ctx := context.Background()

	foodWasm, err := os.ReadFile("../target/wasm32-wasi/debug/foodprovider.wasm")
	if err != nil {
		panic(err)
	}

	// Create a new WebAssembly Runtime.
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx) // This closes everything this Runtime created.
	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	mod, err := r.Instantiate(ctx, foodWasm)
	if err != nil {
		log.Panicf("failed to instantiate module: %v", err)
	}
	createFood := mod.ExportedFunction("create_food")
	results, err := createFood.Call(ctx, 100, 100)
	if err != nil {
		log.Panicf("failed to call: %v", err)
	}
	ptr := results[0] >> 32
	length := results[0] & 0xFFFFFFFF

	data, _ := mod.Memory().Read(uint32(ptr), uint32(length))
	fmt.Println(string(data))
}
