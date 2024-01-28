package hexGrid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHexGrid_GetHex(t *testing.T) {

	grid := CreateHexGrid[int](3)

	got, err := grid.GetHex(0, 0)

	assert.NoError(t, err)
	assert.NotNil(t, got)

	got, err = grid.GetHex(5, 5)
	assert.ErrorContains(t, err, "q not found")
	assert.Nil(t, got)

	got, err = grid.GetHex(0, 5)
	assert.ErrorContains(t, err, "r not found")
	assert.Nil(t, got)

}

func TestHexGrid_GetAllHexes(t *testing.T) {

	hexGrid := CreateHexGrid[int](3)

	got := hexGrid.GetAllHexes()

	assert.Equal(t, 37, len(got))

}

func TestHexGrid_GetNeighbours(t *testing.T) {

	hexGrid := CreateHexGrid[int](2)

	var got hexGroup
	got = hexGrid.GetNeighbours(0, 0)

	assert.True(t, got.contains(+1, 0))
	assert.True(t, got.contains(+1, -1))
	assert.True(t, got.contains(0, -1))
	assert.True(t, got.contains(-1, 0))
	assert.True(t, got.contains(-1, +1))
	assert.True(t, got.contains(0, +1))

	got = hexGrid.GetNeighbours(-2, +2)

	assert.True(t, got.contains(-1, +2))
	assert.True(t, got.contains(-1, +1))
	assert.True(t, got.contains(-2, +1))
	assert.False(t, got.contains(-3, +2))
	assert.False(t, got.contains(-3, +3))
	assert.False(t, got.contains(-2, +3))

}

type hexGroup []*Hex[int]

func (h hexGroup) contains(q, r int) bool {
	for _, v := range h {
		if v.q == q && v.r == r {
			return true
		}
	}
	return false
}
