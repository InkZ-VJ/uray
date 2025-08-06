package uray

import "unique"

func Append[T comparable](arr1 []T, arr2 ...T) []T {
	u := make([]unique.Handle[T], 0, len(arr1)+len(arr2))
	for _, arg := range arr1 {
		u = append(u, unique.Make(arg))
	}

	for _, arg := range arr2 {
		u = append(u, unique.Make(arg))
	}

	out := make([]T, 0, len(u))
	for _, arg := range u {
		out = append(out, arg.Value())
	}

	return out
}
