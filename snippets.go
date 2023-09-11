package main

import "slices"

// Uniq returns a new slice without duplicates
func Uniq(s []int) []int {
	m := mapFrom(s)
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// UniqInPlace overwrites the values in the beginning of the array with the unique ones;
// the re-sizing of the slice means we will ignore the rest of the values because they are duplicates;
// to re-size a given slice, the argument must be a pointer to a descriptor.
// https://go.dev/blog/slices-intro
func UniqInPlace(s *[]int) {
	m := mapFrom(*s)
	*s = (*s)[:len(m):len(m)]
	var i int
	for k := range m {
		(*s)[i] = k
		i++
	}
}

// UniqInPlaceSorted : same as UniqInPlace but Sorted.
func UniqInPlaceSorted(s *[]int) {
	UniqInPlace(s)
	slices.Sort(*s)
}

// mapFrom returns a map
// its keys are the "unique" values we get from the slice argument
// which is useful in this use case;
// the default value associated to keys is simply an empty struct{} instead of a boolean
// because unsafe.Sizeof(struct{}) == 0 and unsafe.Sizeof(boolean) == 1.
func mapFrom(s []int) map[int]struct{} {
	m := make(map[int]struct{})
	for _, e := range s {
		if _, ok := m[e]; ok {
			continue
		}
		m[e] = struct{}{}
	}
	return m
}

