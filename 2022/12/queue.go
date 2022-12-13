package main

type Queue struct {
	data []*Vertex
}

func (q *Queue) Add(v *Vertex) {
	q.data = append(q.data, v)
}

func (q *Queue) Size() int {
	return len(q.data)
}

func (q *Queue) Dequeue() *Vertex {
	v := q.data[0]
	if len(q.data) == 1 {
		q.data = nil // reset slice
	} else {
		q.data = q.data[1:] // drop first value
	}
	return v
}
