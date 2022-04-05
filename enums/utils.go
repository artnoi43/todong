package enums

import "strings"

// toUpper is a generic function for interface enum.
// Since enums are capitalized, it is better to
// call toUpper before comparing values against string enums.

func ToUpper[T ~string](v T) T {
	return T(strings.ToUpper(string(v)))
}
