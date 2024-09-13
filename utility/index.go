package utility

func InArray[T comparable](group []T, target T) bool {
	for _, v := range group {
		if v == target {
			return true
		}
	}
	return false
}
