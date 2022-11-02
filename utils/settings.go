package utils

const debug bool = false
const DefaultOutput string = "dist"

// Hashset for easier lookup
var AcceptedInputFileTypes = map[string]bool{
	".txt": true,
	".md":  true,
}
