package reverse

// Reverse takes a string and returns a reversed string
func Reverse(in string) string {
	reversed := ""
	inRunes := []rune(in)
	for i := len(inRunes) - 1; i >= 0; i-- {
		reversed += string(inRunes[i])
	}
	return reversed
}
