package robotname

import (
	"math/rand"
	"strconv"
	"time"
)

// UsedNames is a list of already taken robot names.
var UsedNames = make(map[string]bool)

// Robot is a simple struct with a robot name
type Robot struct {
	name string
}

// Name returns a random robot name in the form [A-Z][A-Z][100-999]
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	name := generateRandomName()
	for UsedNames[name] == true {
		name = generateRandomName()
	}
	r.name = name
	UsedNames[r.name] = true
	return r.name, nil
}

// Reset resets a given robots name.
func (r *Robot) Reset() *Robot {
	r.name = generateRandomName()
	return r
}

func generateRandomName() string {
	rand.Seed(time.Now().UnixNano())

	letter1 := string((rand.Uint32() % 26) + 65)
	letter2 := string((rand.Uint32() % 26) + 65)
	nums := strconv.Itoa((rand.Int()%899 + 100))

	name := letter1 + letter2 + nums
	return name
}
