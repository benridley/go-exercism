package secret

// Handshake produces a secret handshake based on the binary secret table
func Handshake(n uint) (out []string) {
	if n%2 == 1 {
		out = append(out, "wink")
	}
	if (n>>1)%2 == 1 {
		out = append(out, "double blink")
	}
	if (n>>2)%2 == 1 {
		out = append(out, "close your eyes")
	}
	if (n>>3)%2 == 1 {
		out = append(out, "jump")
	}
	if (n>>4)%2 == 1 {
		out = reverse(out)
	}

	return out
}

func reverse(a []string) []string {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a
}
