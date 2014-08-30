package target

type Target []interface{}

func (t Target) Equals(other Target) bool {
	if len(t) != len(other) {
		return false
	}

	for i, x := range t {
		if x != other[i] {
			return false
		}
	}

	return true
}
