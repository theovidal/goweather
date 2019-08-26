package goweather

// Checks if a specific slice contains a string
func contains(slice []string, text string) bool {
	for _, item := range slice {
		if item == text {
			return true
		}
	}

	return false
}
