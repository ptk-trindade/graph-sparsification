package nodecentrality

func fillSlice[T any](slice []T, val T) {
	for i := range slice {
		slice[i] = val
	}
}