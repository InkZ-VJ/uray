package uray_test

func baseAppendUnique[T comparable](arr1 []T, arr2 ...T) []T {
	uniqueSlice := make([]T, 0, len(arr1)+len(arr2))
	seen := make(map[T]bool)

	for _, element := range arr1 {
		if !seen[element] {
			uniqueSlice = append(uniqueSlice, element)
			seen[element] = true
		}
	}

	for _, element := range arr2 {
		if !seen[element] {
			uniqueSlice = append(uniqueSlice, element)
			seen[element] = true
		}
	}

	return uniqueSlice
}
