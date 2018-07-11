package graph

import (
	"fmt"
	"testing"
)

func TestUndirectedGraph_BFS(t *testing.T) {
	var vertices []int
	type edge struct {
		a, b int
	}
	type args struct {
		vertexCount   int
		edges         []edge
		start         int
		processVertex func(vertexId int, vertexName string, vertexEdges []*Edge)
		processEdge   func(edge *Edge) (terminate bool)
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				vertexCount: 5,
				edges: []edge{
					{0, 1},
					{1, 2},
					{2, 3},
					{3, 4},
				},
				start: 0,
				processVertex: func(id int, name string, edges []*Edge) {
					vertices = append(vertices, id)
				},
				processEdge: nil,
			},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			name: "case 2",
			args: args{
				vertexCount: 5,
				edges: []edge{
					{0, 1},
					{1, 2},
					{2, 3},
					{3, 4},
				},
				start: 4,
				processVertex: func(id int, name string, edges []*Edge) {
					vertices = append(vertices, id)
				},
				processEdge: nil,
			},
			want: []int{4, 3, 2, 1, 0},
		},
		{
			name: "case 3",
			args: args{
				vertexCount: 5,
				edges: []edge{
					{0, 1},
					{1, 2},
					{2, 3},
					{3, 4},
				},
				start: 2,
				processVertex: func(id int, name string, edges []*Edge) {
					vertices = append(vertices, id)
				},
				processEdge: nil,
			},
			want: []int{2, 3, 1, 4, 0},
		},
		{
			name: "case 4",
			args: args{
				vertexCount: 6,
				edges: []edge{
					{0, 1},
					{1, 2},
					{2, 3},
					{3, 4},
					{0, 5},
				},
				start: 0,
				processVertex: func(id int, name string, edges []*Edge) {
					vertices = append(vertices, id)
				},
				processEdge: nil,
			},
			want: []int{0, 5, 1, 2, 3, 4},
		},
		{
			name: "case 5",
			args: args{
				vertexCount: 6,
				edges: []edge{
					{0, 5},
					{0, 4},
					{0, 1},
					{1, 4},
					{1, 2},
					{4, 3},
					{2, 3},
				},
				start: 0,
				processVertex: func(id int, name string, edges []*Edge) {
					vertices = append(vertices, id)
				},
				processEdge: nil,
			},
			want: []int{0, 1, 4, 5, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewUndirectedGraph()
			for i := 0; i < tt.args.vertexCount; i++ {
				g.AddVertex("v for vertex")
			}
			for _, e := range tt.args.edges {
				g.AddEdge(e.a, e.b, 0)
			}
			vertices = make([]int, 0)
			g.BFS(tt.args.start, tt.args.processVertex, tt.args.processEdge)
			if len(tt.want) != len(vertices) {
				t.Fatalf("different expected and actual path want: %v actual: %v", tt.want, vertices)
			}
			for i, j := range tt.want {
				if j != vertices[i] {
					t.Fatalf("different expected and actual path want: %v actual: %v", tt.want, vertices)
				}
			}
		})
	}
}

// Bipartite example. Test if a graph is a bipartite.
// In a graph where we color each vertex with a different color than adjacent vertices,
// the graph is called bipartite if only two colors are enough to complete this operation.
// As an example a binary tree is a bipartite graph. i.e. we can color odd levels with white and even levels with black.
// A graph with 3 vertices A, B and C and 3 edges A - B, B - C and C - A requires 3 colors, therefore is not bipartite.
func ExampleUndirectedGraph_BFS_bipartite() {
	g := NewUndirectedGraph()

	for i := 0; i < 10; i++ {
		g.AddVertex("v for vertex")
	}
	g.AddEdge(0, 1, 0)
	g.AddEdge(0, 2, 0)
	g.AddEdge(1, 3, 0)
	g.AddEdge(1, 4, 0)
	g.AddEdge(2, 5, 0)
	g.AddEdge(2, 6, 0)
	g.AddEdge(3, 7, 0)
	g.AddEdge(3, 8, 0)
	g.AddEdge(4, 9, 0)

	colors := make([]int, 10)

	colors[4] = 1

	bipartite := true

	testEdgeColor := func(edge *Edge) (terminate bool) {
		terminate = false
		if colors[edge.from] > 0 && colors[edge.to] > 0 {
			if colors[edge.from] == colors[edge.to] {
				terminate = true
				bipartite = false
				return
			} else {
				return
			}
		}
		if colors[edge.from] > 0 {
			colors[edge.to] = 3 - colors[edge.from]
		} else if colors[edge.to] > 0 {
			colors[edge.from] = 3 - colors[edge.to]
		}
		return
	}

	// this is a tree, so we expect it to be naturally bipartite
	g.BFS(4, nil, testEdgeColor)

	fmt.Println(bipartite)

	bipartite = true
	colors = make([]int, 5)
	colors[0] = 1
	g2 := NewUndirectedGraph()
	for i := 0; i < 5; i++ {
		g2.AddVertex("v for vertex")
		if i > 0 {
			g2.AddEdge(i, i-1, 0)
		}
	}
	g2.AddEdge(4, 0, 0)

	// g2 is a "pentagon". At least 3 colors are needed.

	g2.BFS(0, nil, testEdgeColor)

	fmt.Println(bipartite)

	// Output:
	// true
	// false

}

// Connected Components example. Test if two vertices in a given graph are connected.
func ExampleUndirectedGraph_BFS_connectedComponents() {
	g := NewUndirectedGraph()
	for i := 0; i < 10; i++ {
		g.AddVertex("v for vertex")
		if i > 0 {
			g.AddEdge(i, i-1, 0)
		}
	}
	connected := false
	testConnectivity := func(testVertex int) func(id int, name string, edges []*Edge) {
		return func(id int, name string, edges []*Edge) {
			if id == testVertex {
				connected = true
			}
		}
	}

	// check if 0 and 9 are connected
	g.BFS(0, testConnectivity(9), nil)

	fmt.Println(connected)
	g.RemoveEdge(3, 4)
	connected = false
	g.BFS(0, testConnectivity(9), nil)

	fmt.Println(connected)
	// Output:
	// true
	// false

}
