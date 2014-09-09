package knnutilities

func Euclidean(row1, row2 []float64, bailout float64) (distance float64) {
	for i := 0; i < len(row1); i++ {
		x := row1[i] - row2[i]
		distance = distance + x*x
		if distance > bailout {
			return bailout
		}
	}

	return
}
