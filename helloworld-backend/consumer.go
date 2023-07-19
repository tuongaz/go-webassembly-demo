package main

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

//go:embed add/target/wasm32-wasi/debug/add.wasm
var wasmAdd []byte

func main() {
	ctx := context.Background()
	module := mustInitWASIRuntime(ctx)

	res, err := module.ExportedFunction("add").Call(ctx, 1, 2)
	if err != nil {
		panic(err)
	}

	fmt.Println(res[0])
}

func mustInitWASIRuntime(ctx context.Context) api.Module {
	r := wazero.NewRuntime(ctx)
	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	module, err := r.Instantiate(ctx, wasmAdd)
	if err != nil {
		panic(err)
	}

	return module
}
