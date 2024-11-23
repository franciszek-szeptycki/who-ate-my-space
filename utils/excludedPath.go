package utils

var excludedPaths = []string{
	"/proc", "/dev",
}

func CheckIsPathExcluded(path string) bool {
	for _, excludedPath := range excludedPaths {
		if path == excludedPath[:len(excludedPath)-1] {
			return true
		}
	}
	return false
}
