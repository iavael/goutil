//go:build go1.18 && !reflect

package goutil

// MemberOfSlice checks if object is element of slice
func MemberOfSlice[K comparable](element K, slice []K) bool {
	for _, v := range slice {
		if element == v {
			return true
		}
	}
	return false
}

// DiffSlices returns differense between two slices
func DiffSlices[K comparable](one []K, two []K) (oldS []K, newS []K) {
	for _, v := range one {
		if !MemberOfSlice(v, two) {
			oldS = append(oldS, v)
		}
	}

	for _, v := range two {
		if !MemberOfSlice(v, one) {
			newS = append(newS, v)
		}
	}

	return
}
