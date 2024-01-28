package hexGrid

type AxialCoords struct {
	q, r int
}

func (a AxialCoords) Q() int {
	return a.q
}

func (a AxialCoords) R() int {
	return a.r
}

type Hex[T interface{}] struct {
	AxialCoords
	content *T
}

func (h *Hex[T]) SetContent(in *T) {
	h.content = in
}

func (h *Hex[T]) GetContent() *T {
	return h.content
}
