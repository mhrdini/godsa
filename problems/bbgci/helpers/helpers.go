package helpers

func GetZeroBasedLetterIndex(c rune) int {
	return int(c - 'a')
}

func SameIntArray(a, b []int) bool {
	for i := range 26 {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func HasAllUnique(s string) bool {
	var mask int

	for _, c := range s {
		bit := 1 << (c - 'a')

		// if bit already set -> duplicate
		if mask&bit != 0 {
			return false
		}

		// set the bit
		mask |= bit
	}

	return true
}
