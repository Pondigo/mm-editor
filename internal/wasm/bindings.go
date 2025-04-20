//go:build js && wasm
// +build js,wasm

package wasm

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/Pondigo/mermaid-editor/internal/core/diagram"
)

// RegisterCallbacks registers all WebAssembly callback functions
func RegisterCallbacks() {
	js.Global().Set("getDiagramAsMermaid", js.FuncOf(getDiagramAsMermaid))
	js.Global().Set("createNewNode", js.FuncOf(createNewNode))
	js.Global().Set("updateNodePosition", js.FuncOf(updateNodePosition))
	js.Global().Set("connectNodes", js.FuncOf(connectNodes))

	fmt.Println("WebAssembly callbacks registered")
}

// getDiagramAsMermaid returns the current diagram as Mermaid syntax
func getDiagramAsMermaid(_ js.Value, _ []js.Value) interface{} {
	if diagram.CurrentDiagram == nil {
		return ""
	}

	return diagram.CurrentDiagram.ToMermaid()
}

// createNewNode creates a new node in the diagram
func createNewNode(_ js.Value, args []js.Value) interface{} {
	if len(args) < 4 {
		return "Error: Missing arguments"
	}

	nodeType := args[0].String()
	text := args[1].String()
	x := args[2].Float()
	y := args[3].Float()

	if diagram.CurrentDiagram == nil {
		diagram.CurrentDiagram = &diagram.Diagram{
			ID: "new-diagram",
		}
	}

	newID := fmt.Sprintf("node%d", len(diagram.CurrentDiagram.Nodes)+1)

	newNode := diagram.Node{
		ID:     newID,
		Type:   nodeType,
		Text:   text,
		X:      x,
		Y:      y,
		Width:  120,
		Height: 60,
	}

	diagram.CurrentDiagram.Nodes = append(diagram.CurrentDiagram.Nodes, newNode)

	// Convert node to JSON to return to JavaScript
	nodeJSON, _ := json.Marshal(newNode)

	return string(nodeJSON)
}

// updateNodePosition updates a node's position
func updateNodePosition(_ js.Value, args []js.Value) interface{} {
	if len(args) < 3 {
		return "Error: Missing arguments"
	}

	nodeID := args[0].String()
	x := args[1].Float()
	y := args[2].Float()

	if diagram.CurrentDiagram == nil {
		return "Error: No diagram exists"
	}

	for i, node := range diagram.CurrentDiagram.Nodes {
		if node.ID == nodeID {
			diagram.CurrentDiagram.Nodes[i].X = x
			diagram.CurrentDiagram.Nodes[i].Y = y
			return "Node position updated"
		}
	}

	return "Error: Node not found"
}

// connectNodes creates a connection between two nodes
func connectNodes(_ js.Value, args []js.Value) interface{} {
	if len(args) < 4 {
		return "Error: Missing arguments"
	}

	sourceID := args[0].String()
	targetID := args[1].String()
	label := args[2].String()
	edgeType := args[3].String()

	if diagram.CurrentDiagram == nil {
		return "Error: No diagram exists"
	}

	newID := fmt.Sprintf("edge%d", len(diagram.CurrentDiagram.Edges)+1)

	newEdge := diagram.Edge{
		ID:       newID,
		Source:   sourceID,
		Target:   targetID,
		Label:    label,
		EdgeType: edgeType,
	}

	diagram.CurrentDiagram.Edges = append(diagram.CurrentDiagram.Edges, newEdge)

	// Convert edge to JSON to return to JavaScript
	edgeJSON, _ := json.Marshal(newEdge)

	return string(edgeJSON)
}
