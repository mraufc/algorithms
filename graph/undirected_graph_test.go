package graph

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

func TestUndirectedGraph_AddVertex(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name  string
		count int
	}{
		{
			name:  "case 1",
			count: 1,
		},
		{
			name:  "case 2",
			count: 100,
		},
		{
			name:  "case 3",
			count: 1000,
		},
		{
			name:  "case 4",
			count: 10000,
		},
		{
			name:  "case 5",
			count: 20000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewUndirectedGraph()
			for i := 0; i < tt.count; i++ {
				name := fmt.Sprintf("V_%v", strconv.Itoa(i+1))
				g.AddVertex(name)
			}
			if len(g.Vertices()) != tt.count {
				log.Fatalln("something went wrong")
			}
		})
	}
}

func TestUndirectedGraph_RemoveVertex(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name        string
		vertexcount int
		deletecount int
	}{
		{
			name:        "case 1",
			vertexcount: 1,
			deletecount: 1,
		},
		{
			name:        "case 2",
			vertexcount: 100,
			deletecount: 20,
		},
		{
			name:        "case 3",
			vertexcount: 1000,
			deletecount: 300,
		},
		{
			name:        "case 4",
			vertexcount: 10000,
			deletecount: 350,
		},
		{
			name:        "case 5",
			vertexcount: 12121,
			deletecount: 2221,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewUndirectedGraph()
			for i := 0; i < tt.vertexcount; i++ {
				name := fmt.Sprintf("V_%v", strconv.Itoa(i+1))
				g.AddVertex(name)
				// add edges from every vertex that will be deleted to every other vertex that has been added so far
				if i > 0 && i < tt.deletecount {
					for j := i - 1; j >= 0; j-- {
						g.AddEdge(i, j, 0)
					}
				}
			}
			if len(g.Vertices()) != tt.vertexcount {
				log.Fatalln("there should be", tt.vertexcount, "vertices")
			}
			if len(g.Edges()) != (tt.deletecount-1)*(tt.deletecount)/2 {
				log.Fatalf("edge count should be %v but is %v \n", (tt.deletecount-1)*(tt.deletecount)/2, len(g.Edges()))
			}
			for i := 0; i < tt.deletecount; i++ {
				g.RemoveVertex(i)
			}
			if len(g.Vertices()) != tt.vertexcount-tt.deletecount {
				log.Fatalln("something went wrong added vertex count", tt.vertexcount, "deleted count", tt.deletecount, "expected count", tt.vertexcount-tt.deletecount)
			}
			if len(g.Edges()) != 0 {
				log.Fatalln("all edges should have been removed")
			}
			// try to remove an out of range vertex
			g.RemoveVertex(tt.vertexcount)
			// try to remove some already removed vertices
			for i := 0; i < tt.deletecount; i++ {
				g.RemoveVertex(i)
			}
			// we expect no changes
			if len(g.Vertices()) != tt.vertexcount-tt.deletecount {
				log.Fatalln("something went wrong added vertex count", tt.vertexcount, "deleted count", tt.deletecount, "expected count", tt.vertexcount-tt.deletecount)
			}
			if len(g.Edges()) != 0 {
				log.Fatalln("all edges should have been removed")
			}
		})
	}
}

func TestUndirectedGraph_AddEdge(t *testing.T) {
	type edge struct {
		v1, v2, w int
	}
	tests := []struct {
		name       string
		vertices   int
		edges      []edge
		testVertex int
		want       int
	}{
		{
			name:     "case 1",
			vertices: 10,
			edges: []edge{
				{0, 1, 1},
				{0, 2, 2},
				{0, 3, 3},
				{5, 0, 10},
			},
			testVertex: 0,
			want:       4,
		},
		{
			name:     "case 2",
			vertices: 100,
			edges: []edge{
				{0, 1, 1},
				{0, 2, 2},
				{5, 0, 10},
				{0, 3, 3},    // 1
				{3, 3, 3},    // should be no-op
				{0, 3, 3},    // should be no-op
				{4, 3, 3},    // 2
				{5, 3, 30},   // 3
				{6, 3, 3},    // 4
				{500, 3, 3},  // should be no-op
				{3, 5000, 3}, // should be no-op
				{3, -1, 3},   // should be no-op
				{-100, 3, 3}, // should be no-op
				{3, 50, 3},   // 5
				{55, 3, 333}, // 6
			},
			testVertex: 3,
			want:       6,
		},
		{
			name:     "case 2",
			vertices: 52,
			edges: []edge{
				{0, 1, 1},
				{0, 2, 2},
				{5, 0, 10},
				{0, 3, 3},
				{3, 3, 3},
				{0, 3, 3},
				{4, 3, 3},
				{5, 3, 30},
				{6, 3, 3},
				{500, 3, 3},
				{3, 5000, 3},
				{3, -1, 3},
				{-100, 3, 3},
				{3, 50, 3},
				{55, 3, 333},
			},
			testVertex: 50,
			want:       1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewUndirectedGraph()
			for i := 0; i < tt.vertices; i++ {
				g.AddVertex(fmt.Sprintf("vertex_%v", strconv.Itoa(i)))
			}
			for i := 0; i < len(tt.edges); i++ {
				g.AddEdge(tt.edges[i].v1, tt.edges[i].v2, tt.edges[i].w)
			}
			edges := g.VertexEdges(tt.testVertex)
			if len(edges) != tt.want {
				t.Fatalf("invalid number of edges, want: %v, actual: %v", tt.want, len(edges))
			}
			for _, ee := range edges {
				for _, e := range tt.edges {
					if (ee.from == e.v1 && ee.to == e.v2) || (ee.to == e.v1 && ee.from == e.v2) {
						if ee.weight != e.w {
							t.Fatalf("invalid weight, want: %v, actual: %v", e.w, ee.weight)
						}
					}
				}
			}
		})
	}
}

func TestUndirectedGraph_RemoveEdge(t *testing.T) {
	type edge struct {
		v1, v2, w int
	}
	type removeEdge struct {
		v1, v2 int
	}
	tests := []struct {
		name        string
		vertices    int
		addEdges    []edge
		removeEdges []removeEdge
		testVertex  int
		want        int
	}{
		{
			name:     "case 1",
			vertices: 10,
			addEdges: []edge{
				{0, 1, 1},
				{0, 2, 2},
				{0, 3, 3},
				{5, 0, 10},
			},
			removeEdges: []removeEdge{
				{0, 2},
			},
			testVertex: 0,
			want:       3,
		},
		{
			name:     "case 2",
			vertices: 100,
			addEdges: []edge{
				{0, 1, 1},
				{0, 2, 2},
				{5, 0, 10},
				{0, 3, 3},    // 1
				{3, 3, 3},    // should be no-op
				{0, 3, 3},    // should be no-op
				{4, 3, 3},    // 2
				{5, 3, 30},   // 3
				{6, 3, 3},    // 4
				{500, 3, 3},  // should be no-op
				{3, 5000, 3}, // should be no-op
				{3, -1, 3},   // should be no-op
				{-100, 3, 3}, // should be no-op
				{3, 50, 3},   // 5
				{55, 3, 333}, // 6
			},
			removeEdges: []removeEdge{
				{3, 0},
				{3, 0},           // should be no-op
				{3, -1},          // should be no-op
				{-1, 3},          // should be no-op
				{3, 10000},       // should be no-op
				{1000000, 10000}, // should be no-op
				{3, 4},
				{3, 6},
				{50, 3},
				{3, 55},
				{5, 3},
			},
			testVertex: 3,
			want:       0,
		},
		{
			name:     "case 2",
			vertices: 52,
			addEdges: []edge{
				{0, 1, 1},
				{0, 2, 2},
				{5, 0, 10},
				{0, 3, 3},
				{3, 3, 3},
				{0, 3, 3},
				{4, 3, 3},
				{5, 3, 30},
				{6, 3, 3},
				{500, 3, 3},
				{3, 5000, 3},
				{3, -1, 3},
				{-100, 3, 3},
				{3, 50, 3}, // 1
				{5, 50, 3}, // 2
				{50, 6, 3}, // 3
				{55, 3, 333},
			},
			removeEdges: []removeEdge{
				{6, 50},
			},
			testVertex: 50,
			want:       2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewUndirectedGraph()
			for i := 0; i < tt.vertices; i++ {
				g.AddVertex(fmt.Sprintf("vertex_%v", strconv.Itoa(i)))
			}
			for i := 0; i < len(tt.addEdges); i++ {
				g.AddEdge(tt.addEdges[i].v1, tt.addEdges[i].v2, tt.addEdges[i].w)
			}
			for i := 0; i < len(tt.removeEdges); i++ {
				g.RemoveEdge(tt.removeEdges[i].v1, tt.removeEdges[i].v2)
			}
			edges := g.VertexEdges(tt.testVertex)
			if len(edges) != tt.want {
				t.Fatalf("invalid number of edges, want: %v, actual: %v", tt.want, len(edges))
			}
			for _, ee := range edges {
				for _, e := range tt.addEdges {
					if (ee.from == e.v1 && ee.to == e.v2) || (ee.to == e.v1 && ee.from == e.v2) {
						if ee.weight != e.w {
							t.Fatalf("invalid weight, want: %v, actual: %v", e.w, ee.weight)
						}
					}
				}
			}
		})
	}
}
