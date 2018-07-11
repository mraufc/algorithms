package graph

// queue is a simple queue implementation using a linked list
type queue struct {
	head *node
}
type node struct {
	value int
	next  *node
}

// BFS performs a breadth first search in the graph g starting from a vertex whose id is given as input.
// processVertex function is called with vertex id, vertex name and all the 'outgoing' edges of the vertex once a vertex is encountered
// processEdge function is called once an edge is traversed
func (g *UndirectedGraph) BFS(start int, processVertex func(vertexId int, vertexName string, vertexEdges []*Edge), processEdge func(edge *Edge) (terminate bool)) {
	g.mtx.RLock()
	defer g.mtx.RUnlock()
	// check start validity
	if start < 0 || start >= len(g.vertices) {
		return
	}
	if g.vertices[start] == nil {
		return
	}

	q := newQueue()
	var v, y, status int
	var e *edge
	var t bool
	discovered := make([]bool, len(g.vertices))
	processed := make([]bool, len(g.vertices))
	parent := make([]int, len(g.vertices))
	q.enqueue(start)

	discovered[start] = true
	for {
		v, status = q.dequeue()
		if status == -1 {
			break
		}
		if g.vertices[v] == nil {
			continue
		}
		if processVertex != nil {
			processVertex(g.vertices[v].id, g.vertices[v].name, g.vertexEdgesSafe(v))
		}
		processed[v] = true

		e = g.vertices[v].edges
		for e != nil {
			y = e.to
			if !processed[y] {
				if processEdge != nil {
					if t = processEdge(&Edge{from: v, to: y, weight: e.weight}); t {
						return
					}
				}
			}
			if !discovered[y] {
				q.enqueue(y)
				discovered[y] = true
				parent[y] = v
			}
			e = e.next
		}
	}
}

func newQueue() *queue {
	return new(queue)
}

func (q *queue) enqueue(v int) {
	n := &node{value: v}

	if q.head == nil {
		q.head = n
		return
	}
	p := q.head

	for p.next != nil {
		p = p.next
	}
	p.next = n
}

// status is -1 if there is nothing to dequeue
func (q *queue) dequeue() (value int, status int) {

	if q.head == nil {
		status = -1
		return
	}
	value = q.head.value
	q.head = q.head.next
	return
}
