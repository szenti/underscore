package underscore

// Map produces a new slice of values by mapping each value in the slice through
// a transform function.
func Map[T, P any](values []T, transform func(T) P) []P {
	res := make([]P, len(values))
	for i := range values {
		res[i] = transform(values[i])
	}
	return res
}
