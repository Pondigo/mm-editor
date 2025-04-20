//go:build js && wasm
// +build js,wasm

package diagram

import (
	"fmt"
	"syscall/js"
)

// Node represents a diagram node
type Node struct {
	ID     string
	Type   string
	Text   string
	X      float64
	Y      float64
	Width  float64
	Height float64
}

// Diagram represents a mermaid diagram
type Diagram struct {
	ID    string
	Nodes []Node
	Edges []Edge
}

// Edge represents a connection between nodes
type Edge struct {
	ID       string
	Source   string
	Target   string
	Label    string
	EdgeType string
}

// CurrentDiagram is the active diagram being edited
var CurrentDiagram *Diagram

// CreateTestDiagram creates a simple test diagram
func CreateTestDiagram() {
	CurrentDiagram = &Diagram{
		ID: "test-diagram",
		Nodes: []Node{
			{
				ID:     "node1",
				Type:   "process",
				Text:   "Process A",
				X:      100,
				Y:      100,
				Width:  120,
				Height: 60,
			},
			{
				ID:     "node2",
				Type:   "process",
				Text:   "Process B",
				X:      300,
				Y:      200,
				Width:  120,
				Height: 60,
			},
		},
		Edges: []Edge{
			{
				ID:       "edge1",
				Source:   "node1",
				Target:   "node2",
				Label:    "connects to",
				EdgeType: "arrow",
			},
		},
	}

	// Log the test diagram creation
	fmt.Println("Test diagram created with", len(CurrentDiagram.Nodes), "nodes and",
		len(CurrentDiagram.Edges), "edges")

	// Notify JavaScript that the diagram is ready
	js.Global().Call("diagramReady", CurrentDiagram.ID)
}

// ToMermaid converts the diagram to Mermaid syntax
func (d *Diagram) ToMermaid() string {
	mermaid := "graph TD\n"

	// Add nodes
	for _, node := range d.Nodes {
		mermaid += fmt.Sprintf("    %s[%s]\n", node.ID, node.Text)
	}

	// Add edges
	for _, edge := range d.Edges {
		edgeSymbol := "-->"
		if edge.EdgeType == "dotted" {
			edgeSymbol = "-.->"
		}

		if edge.Label != "" {
			mermaid += fmt.Sprintf("    %s %s|%s| %s\n", edge.Source, edgeSymbol, edge.Label, edge.Target)
		} else {
			mermaid += fmt.Sprintf("    %s %s %s\n", edge.Source, edgeSymbol, edge.Target)
		}
	}

	return mermaid
}
