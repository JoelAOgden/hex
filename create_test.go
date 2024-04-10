package hex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateHexGrid_0(t *testing.T) {

	got := CreateHexGrid[int](0)

	assert.Equal(t, uint(0), got.radius)

	assert.Equal(t, 1, len(got.hexes))

	assert.Equal(t, 1, len(got.hexes[0]))
}

func TestCreateHexGrid_1(t *testing.T) {

	got := CreateHexGrid[int](1)

	assert.Equal(t, uint(1), got.radius)

	assert.Equal(t, 3, len(got.hexes))

	assert.Equal(t, 2, len(got.hexes[-1]))
	assert.Equal(t, 3, len(got.hexes[0]))
	assert.Equal(t, 2, len(got.hexes[1]))
}

func TestCreateHexGrid_2(t *testing.T) {

	got := CreateHexGrid[int](2)

	assert.Equal(t, uint(2), got.radius)

	assert.Equal(t, 5, len(got.hexes))

	assert.Equal(t, 3, len(got.hexes[-2]))
	assert.Equal(t, 4, len(got.hexes[-1]))
	assert.Equal(t, 5, len(got.hexes[0]))
	assert.Equal(t, 4, len(got.hexes[1]))
	assert.Equal(t, 3, len(got.hexes[2]))
}

func TestCreateHexGrid_3(t *testing.T) {

	got := CreateHexGrid[int](3)

	assert.Equal(t, uint(3), got.radius)

	assert.Equal(t, 7, len(got.hexes))

	assert.Equal(t, 4, len(got.hexes[-3]))
	assert.Equal(t, 5, len(got.hexes[-2]))
	assert.Equal(t, 6, len(got.hexes[-1]))
	assert.Equal(t, 7, len(got.hexes[0]))
	assert.Equal(t, 6, len(got.hexes[1]))
	assert.Equal(t, 5, len(got.hexes[2]))
	assert.Equal(t, 4, len(got.hexes[3]))
}
