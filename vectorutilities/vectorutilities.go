package vectorutilities

func Add(x, y []float64) []float64 {
    r := make([]float64, len(x))
    for i := range r {
        r[i] = x[i] + y[i]
    }
    return r
}

func Scale(a float64, x []float64) []float64 {
    r := make([]float64, len(x))
    for i := range r {
        r[i] = a * x[i]
    }
    return r
}
