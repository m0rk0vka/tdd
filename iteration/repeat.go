package iteration

// Repeat take characters and repeated count
// and return string of characters repeted
// input amount times.
func Repeat(characters string, repeatedCount int) string {
	var repeated string
	for i := 0; i < repeatedCount; i++ {
		repeated += characters
	}
	return repeated
}
