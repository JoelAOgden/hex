package hex

import "math"

func CreateHexGrid[T interface{}](radius uint) *HexGrid[T] {

	r := int(radius)
	hexes := make(map[int]map[int]*Hex[T])

	for q := -r; q <= r; q++ {

		hexes[q] = make(map[int]*Hex[T])

		r1 := int(math.Max(float64(-r), float64(-q-r)))
		r2 := int(math.Min(float64(r), float64(-q+r)))

		for r := r1; r <= r2; r++ {
			hexes[q][r] = &Hex[T]{
				AxialCoords: AxialCoords{
					q: q,
					r: r,
				},
			}
		}
	}

	return &HexGrid[T]{
		hexes:  hexes,
		radius: radius,
	}

}
