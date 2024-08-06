package functions

func charValidation(str string) bool {
	for _, char := range str {
		if (char < 32 || char > 126) && char != '\n' {
			return false
		}
	}
	return true
}
