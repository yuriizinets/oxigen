package main

import "github.com/kyoto-framework/kyoto/v2"

// PIndexState is a state of PIndex.
type PUIState struct {
	Generator *kyoto.ComponentF[CGeneratorState]
}

// PUI is a generator UI page.
func PUI(ctx *kyoto.Context) (state PUIState) {
	// Setup rendering
	kyoto.Template(ctx, "page.ui.go.html")
	// Init components
	state.Generator = kyoto.Use(ctx, CGenerator(&CGeneratorArgs{}))
	// Return
	return
}
