package graph

import (
	"sync"
)

// UndirectedGraph is adjacency list graph structure
type UndirectedGraph struct {
	mtx      *sync.RWMutex
	vertices []*vertex
}

// Vertex is used for export purposes
type Vertex struct {
	id   int
	name string
}
type vertex struct {
	id    int    // unique identifier
	name  string // name of the vertex - not necessarily needed but anyway
	edges *edge  // 'outgoing' edges
}

// Edge is used for export purposes
type Edge struct {
	from, to, weight int
}
type edge struct {
	to     int //  vertex id
	weight int
	next   *edge
}

// NewUndirectedGraph returns a new undirected graph
func NewUndirectedGraph() *UndirectedGraph {
	return &UndirectedGraph{
		mtx:      new(sync.RWMutex),
		vertices: make([]*vertex, 0),
	}
}

// AddVertex adds a named vertex to the graph. Graph's counter is used as an id for the vertex.
// returns the id of the added vertex.
func (g *UndirectedGraph) AddVertex(name string) int {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	v := &vertex{
		id:    len(g.vertices),
		name:  name,
		edges: nil,
	}
	g.vertices = append(g.vertices, v)
	return v.id
}

// RemoveVertex removes a vertex and all associated edges
func (g *UndirectedGraph) RemoveVertex(a int) {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	if a >= len(g.vertices) || a < 0 {
		return
	}
	if g.vertices[a] == nil {
		return
	}
	e := g.vertices[a].edges
	q := make([]int, 0)
	for e != nil {
		q = append(q, e.to)
		e = e.next
	}
	for _, b := range q {
		if g.vertices[b] != nil {
			if a == b {
				continue
			}
			g.removeEdgeSafe(a, b)
		}
	}
	g.vertices[a] = nil
}

// Vertices returns all the vertex id and names in the graph, returns nil if the graph has no vertices
func (g *UndirectedGraph) Vertices() []*Vertex {
	g.mtx.RLock()
	defer g.mtx.RUnlock()
	r := make([]*Vertex, 0)
	for _, v := range g.vertices {
		if v == nil {
			continue
		}
		r = append(r, &Vertex{v.id, v.name})
	}
	if len(r) == 0 {
		r = nil
	}
	return r
}

// VertexEdges lists all the edges that are connected to a given vertex
func (g *UndirectedGraph) VertexEdges(a int) []*Edge {
	g.mtx.RLock()
	defer g.mtx.RUnlock()

	if a >= len(g.vertices) || a < 0 {
		return nil
	}

	if g.vertices[a] == nil {
		return nil
	}
	return g.vertexEdgesSafe(a)
}

func (g *UndirectedGraph) vertexEdgesSafe(a int) []*Edge {
	r := make([]*Edge, 0)
	e := g.vertices[a].edges

	for e != nil {
		r = append(r, &Edge{g.vertices[a].id, e.to, e.weight})
		e = e.next
	}

	if len(r) == 0 {
		r = nil
	}
	return r
}

// AddEdge adds an edge with weight w between vertices with id a and b.
// updates the weight if an edge already exists.
func (g *UndirectedGraph) AddEdge(a, b, w int) {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	if a >= len(g.vertices) || b >= len(g.vertices) || a < 0 || b < 0 {
		return
	}
	if g.vertices[a] == nil || g.vertices[b] == nil {
		return
	}
	if a == b {
		return
	}
	a2bexists, b2aexists := false, false
	var e *edge
	e = g.vertices[a].edges
	for e != nil {
		if e.to == b {
			e.weight = w
			a2bexists = true
			break
		}
		e = e.next
	}
	if !a2bexists {
		a2b := &edge{
			to:     b,
			weight: w,
			next:   g.vertices[a].edges,
		}
		g.vertices[a].edges = a2b
	}
	e = g.vertices[b].edges
	for e != nil {
		if e.to == a {
			e.weight = w
			b2aexists = true
			break
		}
		e = e.next
	}
	if !b2aexists {
		b2a := &edge{
			to:     a,
			weight: w,
			next:   g.vertices[b].edges,
		}
		g.vertices[b].edges = b2a
	}
}

// RemoveEdge removes an edge between two vertices with ids a and b.
func (g *UndirectedGraph) RemoveEdge(a, b int) {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	if a >= len(g.vertices) || b >= len(g.vertices) || a < 0 || b < 0 {
		return
	}
	if g.vertices[a] == nil || g.vertices[b] == nil {
		return
	}
	if a == b {
		return
	}
	g.removeEdgeSafe(a, b)
}

// Edges lists all the edges in the graph. Returns nil if there are no edges
func (g *UndirectedGraph) Edges() []*Edge {
	g.mtx.RLock()
	defer g.mtx.RUnlock()
	r := make([]*Edge, 0)
	var e *edge
	for _, v := range g.vertices {
		if v == nil {
			continue
		}
		e = v.edges
		for e != nil {
			// there are two 'edges' per edge in the structure, i.e 2 to 3 and 3 to 2
			// just display 2 to 3
			if e.to > v.id {
				r = append(r, &Edge{v.id, e.to, e.weight})
			}
			e = e.next
		}
	}

	if len(r) == 0 {
		r = nil
	}
	return r

}

// removeEdgeSafe is called to remove an edge between a and b.
// it's assumed that validity of vertices a and b are checked before calling this function
func (g *UndirectedGraph) removeEdgeSafe(a, b int) {
	var e, prev *edge
	e = g.vertices[a].edges
	for e != nil {
		if e.to == b {
			if prev != nil {
				prev.next = e.next
			} else {
				g.vertices[a].edges = e.next
			}
			break
		}
		prev = e
		e = e.next
	}
	e = g.vertices[b].edges
	prev = nil
	for e != nil {
		if e.to == a {
			if prev != nil {
				prev.next = e.next
			} else {
				g.vertices[b].edges = e.next
			}
			break
		}
		prev = e
		e = e.next
	}
}
