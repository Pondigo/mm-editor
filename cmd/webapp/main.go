//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"

	"github.com/Pondigo/mermaid-editor/internal/core/diagram"
	"github.com/Pondigo/mermaid-editor/internal/wasm"
)

func main() {
	fmt.Println("Mermaid Editor WebAssembly initialized")

	// Create a channel to keep the program running
	c := make(chan struct{}, 0)

	// Register JavaScript callbacks
	wasm.RegisterCallbacks()

	// Create test diagram for development
	diagram.CreateTestDiagram()

	// Keep the program running
	<-c
}
