package hexGrid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHex_SetContent(t *testing.T) {

	hex := Hex[int]{
		AxialCoords: AxialCoords{
			q: 0,
			r: 0,
		},
	}

	assert.Nil(t, hex.content)

	in := 1
	hex.SetContent(&in)

	assert.Equal(t, &in, hex.content)

}

func TestHex_GetContent(t *testing.T) {

	in := 1
	hex := Hex[int]{
		AxialCoords: AxialCoords{
			q: 0,
			r: 0,
		},
		content: &in,
	}

	got := hex.GetContent()

	assert.Equal(t, &in, got)
}
