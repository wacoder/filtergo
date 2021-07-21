package filtergo

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func newGHFilter() *GHFilter {
	ghFilter := GHFilter{
		x:  0,
		dx: 0,
		dt: 1,
		g:  0.6,
		h:  0.02,
	}
	return &ghFilter
}

func TestGHFilter(t *testing.T) {
	gh := newGHFilter()

	// measurement vector
	for i := 0.1; i < 100; i++ {
		z := rand.Float64() + float64(i)
		x, dx := gh.Update(z)

		if math.Abs(x-i) > 1 {
			fmt.Println(z, x, dx)
			t.Error("Prediction is off")
		}
	}
}
