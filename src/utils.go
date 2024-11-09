package cfgo

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}

	return v
}

func joinMaps[T any](mapsIter ...map[string]T) map[string]T {
	outMap := map[string]T{}

	for _, candidateMap := range mapsIter {
		for k, v := range candidateMap {
			_, ok := outMap[k]

			if ok {
				continue
			}

			outMap[k] = v
		}
	}

	return outMap
}

func sliceOrDefault[T any](slice []T, defaults []T) []T {
	if len(slice) == 0 {
		return defaults
	}

	return slice
}

func sliceHas[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}
