package hex

import "fmt"

type HexGrid[T interface{}] struct {
	hexes  map[int]map[int]*Hex[T]
	radius uint
}

func (h *HexGrid[T]) GetHex(q, r int) (*Hex[T], error) {

	if _, ok := h.hexes[q]; !ok {
		return nil, fmt.Errorf("q not found")
	}

	if _, ok := h.hexes[q][r]; !ok {
		return nil, fmt.Errorf("r not found")
	}

	return h.hexes[q][r], nil
}

func (h *HexGrid[T]) GetAllHexes() []*Hex[T] {
	out := make([]*Hex[T], 0)

	for _, v := range h.hexes {
		for _, hex := range v {
			out = append(out, hex)
		}
	}

	return out

}

func (h *HexGrid[T]) GetNeighbours(q, r int) []*Hex[T] {

	coords := []struct {
		qOffset int
		rOffset int
	}{
		{+1, 0}, {+1, -1}, {0, -1},
		{-1, 0}, {-1, +1}, {0, +1},
	}

	out := make([]*Hex[T], 0)
	for _, v := range coords {
		hex, _ := h.GetHex(q+v.qOffset, r+v.rOffset)
		if hex == nil {
			continue
		}

		out = append(out, hex)
	}

	return out

}

func (h *HexGrid[T]) GetRadius() uint { return h.radius }
