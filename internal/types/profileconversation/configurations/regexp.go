package configurations

type Regexp struct {
	// FindInSubMatches is a boolean value that indicates whether or not to find matches in submatches.
	// The `Submatch` variants include information about
	// both the whole-pattern matches and the submatches
	// within those matches. For example this will return
	FindInSubmatches bool
}
