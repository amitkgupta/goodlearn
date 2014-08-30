package knnutilities

import (
	"math"
)

func Euclidean(row1, row2 []float64, bailout float64) (distance float64) {
	for i := 0; i < len(row1); i++ {
		distance = distance + math.Pow(row1[i]-row2[i], 2)
		if distance > bailout {
			break
		}
	}

	return
}
