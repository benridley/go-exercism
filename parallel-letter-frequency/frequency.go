package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in an array of texts concurrently
// and returns the data as a FreqMap.
func ConcurrentFrequency(strings []string) FreqMap {
	textChannel := make(chan FreqMap)
	for _, text := range strings {
		go func(text string) {
			textChannel <- Frequency(text)
		}(text)
	}

	combinedMap := FreqMap{}
	for i := 0; i < len(strings); i++ {
		for key, value := range <-textChannel {
			combinedMap[key] += value
		}
	}

	return combinedMap
}
